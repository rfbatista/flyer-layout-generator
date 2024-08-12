package ports

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type Controller interface {
	Load(e *echo.Echo) error
}

func AsController(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Controller)),
		fx.ResultTags(`group:"controller"`),
	)
}
