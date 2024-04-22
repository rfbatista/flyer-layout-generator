package infra

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"algvisual/shared"
)

type HTTPServerParams struct {
	fx.In
	Logger      *zap.Logger
	Config      *AppConfig
	Controllers []shared.Controller `group:"controller"`
}

func customHTTPErrorHandler(err error, c echo.Context) {
	c.Logger().Error(err)
	he, _ := err.(*echo.HTTPError)
	he = &echo.HTTPError{
		Code:    http.StatusBadRequest,
		Message: err.Error(),
	}
	code := he.Code
	message := he.Message
	if _, ok := he.Message.(string); ok {
		message = map[string]interface{}{"message": err.Error()}
	}
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(he.Code)
		} else {
			err = c.JSON(code, message)
		}
		if err != nil {
			c.Echo().Logger.Error(err)
		}
	}
}

func NewHTTPServer(p HTTPServerParams) *echo.Echo {
	e := echo.New()
	e.HTTPErrorHandler = customHTTPErrorHandler
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
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Static("/dist", p.Config.DistFolderPath)
	for _, controller := range p.Controllers {
		err := controller.Load(e)
		if err != nil {
			p.Logger.Error("failed to load controller", zap.Error(err))
		}
	}
	for _, r := range e.Routes() {
		p.Logger.Info(fmt.Sprintf("%s\t\t%s", r.Method, r.Path))
	}
	return e
}
