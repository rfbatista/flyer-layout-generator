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
	protected(NewUploadDesignAPI),
	protected(NewListDesignElementsAPI),
	public(NewLoginAPI),
	protected(NewListTemplatesAPI),
	protected(NewCreateTemplateAPI),
	protected(NewSetPhotoshopBackgroundAPI),
	protected(NewListDesignFilesAPI),
	protected(NewListGeneratedImagesAPI),
	protected(NewCreateComponentAPI),
	protected(NewGetDesignByIDAPI),
	protected(NewGenerateDesignAPI),
	protected(NewRemoveComponentAPI),
	protected(NewListComponentsByFileIDAPI),
	protected(NewUploadImage),
	protected(NewDownloadImage),
	protected(NewPageCreateTemplate),
	protected(NewPageHome),
	protected(NewWebUploadDesignAPI),
	protected(NewPageDefineElements),
	protected(NewPageProccessDesign),
	protected(NewWebProccessDesign),
	protected(NewDownloadDesignFiles),
))
