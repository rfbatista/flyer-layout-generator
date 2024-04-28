package api

import (
	"fmt"

	"github.com/rfbatista/apitools"
	"go.uber.org/fx"

	"algvisual/internal/ports"
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
		fx.As(new(ports.Controller)),
		fx.ResultTags(`group:"controller"`),
	)
}

var Module = fx.Options(fx.Provide(
	AsController(NewWebController),
	protected(NewUploadPhotoshopAPI),
	protected(NewListPhotoshopElementsAPI),
	public(NewLoginAPI),
	protected(NewListTemplatesAPI),
	protected(NewCreateTemplateAPI),
	protected(NewSetPhotoshopBackgroundAPI),
	protected(NewListPhotoshopFilesAPI),
	protected(NewListGeneratedImagesAPI),
	protected(NewCreateComponentAPI),
	protected(NewGetPhotoshopByIDAPI),
))
