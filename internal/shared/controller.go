package shared

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func AsController(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Controller)),
		fx.ResultTags(`group:"controller"`),
	)
}

type Controller interface {
	Load(e *echo.Echo) error
}
