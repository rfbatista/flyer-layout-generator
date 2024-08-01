package infra

import (
	"algvisual/internal/ports"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type HTTPServerParams struct {
	fx.In
	Logger      *zap.Logger
	Config      *AppConfig
	Controllers []ports.Controller `group:"controller"`
	Cognito     *Cognito
	Pool        *pgxpool.Pool
	Sse         *ServerSideEventManager
}

type HTTPError struct {
	Code       string    `json:"code,omitempty"`
	Message    string    `json:"message,omitempty"`
	Details    string    `json:"details,omitempty"`
	Timestamp  time.Time `json:"timestamp,omitempty"`
	Path       string    `json:"path,omitempty"`
	Suggestion string    `json:"suggestion,omitempty"`
}

type HTTPErrorResult struct {
	Status      string    `json:"status,omitempty"`
	StatusCode  int       `json:"status_code,omitempty"`
	RequestID   string    `json:"request_id,omitempty"`
	DocumentURL string    `json:"document_url,omitempty"`
	Error       HTTPError `json:"error,omitempty"`
}

type APIHealth struct {
	Status string `json:"status,omitempty"`
	Error  string `json:"error,omitempty"`
	Env    string `json:"env,omitempty"`
}

func NewHTTPServer(p HTTPServerParams) *echo.Echo {
	e := echo.New()
	e.HTTPErrorHandler = HTTPErrorHandler
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(
		middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
			LogURI:      true,
			LogStatus:   true,
			LogError:    true,
			HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
			LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
				p.Logger.Info(
					"request",
					zap.String("method", v.Method),
					zap.String("URI", v.URI),
					zap.Int("status", v.Status),
				)
				return nil
			},
		}))
	e.Use(middleware.Secure())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("page", c.Request().URL)
			return next(c)
		}
	})

	e.GET("/sse", p.Sse.HandleConnection)
	SetupStaticServer(p, e)

	e.GET("/api/health", func(c echo.Context) error {
		err := p.Pool.Ping(c.Request().Context())
		if err != nil {
			status := APIHealth{
				Status: "error",
				Error:  err.Error(),
				Env:    p.Config.APPENV,
			}
			return c.JSONPretty(http.StatusOK, status, "")
		} else {
			status := APIHealth{
				Status: "ok",
				Env:    p.Config.APPENV,
			}
			return c.JSONPretty(http.StatusOK, status, "")
		}
	})
	if p.Config.APPENV == "prod" {
		e.Use(NewAuthMiddleware(p.Cognito))
	}
	for _, controller := range p.Controllers {
		err := controller.Load(e)
		if err != nil {
			p.Logger.Error("failed to load controller", zap.Error(err))
		}
	}
	for _, r := range e.Routes() {
		p.Logger.Info(fmt.Sprintf("%s\t%s", r.Method, r.Path))
	}
	e.HideBanner = true
	return e
}
