package infra

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"algvisual/internal/ports"
	"algvisual/internal/shared"
)

type HTTPServerParams struct {
	fx.In
	Logger      *zap.Logger
	Config      *AppConfig
	Controllers []ports.Controller `group:"controller"`
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

func customHTTPErrorHandler(err error, c echo.Context) {
	var result HTTPErrorResult
	var errorDetails HTTPError
	he, ok := err.(*shared.AppError)
	result.Status = "error"
	if ok {
		result.StatusCode = he.StatusCode
		errorDetails.Message = he.Message
		errorDetails.Timestamp = he.Timestamp
		errorDetails.Details = he.Detail
		result.Error = errorDetails
	} else {
		result.StatusCode = 500
		errorDetails.Message = err.Error()
		errorDetails.Timestamp = time.Now()
	}
	result.Error = errorDetails

	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(result.StatusCode)
		} else {
			err = c.JSON(result.StatusCode, result)
		}
		if err != nil {
			c.Echo().Logger.Error(err)
		}
	}
}

func NewHTTPServer(p HTTPServerParams) *echo.Echo {
	e := echo.New()
	e.HTTPErrorHandler = customHTTPErrorHandler
	// e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
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
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       p.Config.AssetsFolderPath,
		Index:      "index.html",
		Browse:     false,
		HTML5:      true,
		IgnoreBase: false,
		Filesystem: nil,
	}))
	for _, controller := range p.Controllers {
		err := controller.Load(e)
		if err != nil {
			p.Logger.Error("failed to load controller", zap.Error(err))
		}
	}
	for _, r := range e.Routes() {
		p.Logger.Info(fmt.Sprintf("%s\t%s", r.Method, r.Path))
	}
	return e
}
