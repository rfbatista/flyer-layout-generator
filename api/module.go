package api

import (
	"fmt"

	"github.com/rfbatista/apitools"
	"go.uber.org/fx"

	"algvisual/shared"
)

func protected(f any) any {
	return AsRoute(f, "protected")
}

func public(f any) any {
	return AsRoute(f, "public")
}

func AsRoute(f any, name string) any {
	return fx.Annotate(
		f,
		fx.As(new(apitools.Handler)),
		fx.ResultTags(fmt.Sprintf(`group:"%s"`, name)),
	)
}

func AsController(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(shared.Controller)),
		fx.ResultTags(`group:"controller"`),
	)
}

var Module = fx.Options(fx.Provide(
	AsController(NewWebController),
	protected(NewSavePhotoshopAPI),
	protected(NewListPhotoshop),
	public(NewLoginAPI),
))
