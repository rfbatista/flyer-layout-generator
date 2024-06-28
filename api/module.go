package api

import (
	"algvisual/internal/ports"
	"fmt"

	"github.com/rfbatista/apitools"
	"go.uber.org/fx"
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
	AsController(NewProjectsController),
	AsController(NewClientsController),
	AsController(NewAdvertiserController),
	AsController(NewDesignController),
	AsController(NewAssetsController),
	public(NewLoginAPI),
	protected(NewListTemplatesAPI),
	protected(NewCreateTemplateAPI),
	protected(NewSetPhotoshopBackgroundAPI),
	protected(NewListGeneratedImagesAPI),
	protected(NewRemoveComponentAPI),
	protected(NewListComponentsByFileIDAPI),
	protected(NewUploadImage),
	protected(NewDownloadImage),
	protected(NewDownloadDesignFiles),
))
