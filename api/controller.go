package api

import (
	"github.com/labstack/echo/v4"
	"github.com/rfbatista/apitools"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"algvisual/internal/ports"
)

type WebControllerParams struct {
	fx.In
	Logger    *zap.Logger
	Protected []apitools.Handler `group:"protected"`
	Public    []apitools.Handler `group:"public"`
}

func NewWebController(p WebControllerParams) (ports.Controller, error) {
	return WebController{
		Logger:    p.Logger,
		Protected: p.Protected,
		Public:    p.Public,
	}, nil
}

type WebController struct {
	Logger    *zap.Logger
	Protected []apitools.Handler
	Public    []apitools.Handler
}

func (p WebController) Load(e *echo.Echo) error {
	p.loadHandlers(e, p.Public)
	p.loadHandlers(e, p.Protected)
	return nil
}

func (p WebController) loadHandlers(e *echo.Echo, apis []apitools.Handler) {
	for _, handler := range apis {
		if handler == nil {
			continue
		}
		p.Logger.Info(
			"loading handler",
			zap.String("method", handler.Method().String()),
			zap.String("path", handler.Path()),
		)
		switch handler.Method() {
		case apitools.DELETE:
			e.DELETE(handler.Path(), handler.Handle(), handler.Middleware()...)
		case apitools.GET:
			e.GET(handler.Path(), handler.Handle(), handler.Middleware()...)
		case apitools.PUT:
			e.PUT(handler.Path(), handler.Handle(), handler.Middleware()...)
		case apitools.POST:
			e.POST(handler.Path(), handler.Handle(), handler.Middleware()...)
		case apitools.PATCH:
			e.PATCH(handler.Path(), handler.Handle(), handler.Middleware()...)
		}
	}
}
