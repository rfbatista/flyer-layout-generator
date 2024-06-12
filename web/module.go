package web

import (
	"algvisual/internal/ports"
	"algvisual/web/views/components"
	"algvisual/web/views/files"
	"algvisual/web/views/home"
	"algvisual/web/views/jobs"
	"algvisual/web/views/templates"
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
	protected(home.NewPageHome),
	protected(home.CreateImage),
	protected(files.NewPage),
	protected(files.NewUploadDesignAPI),
	protected(files.NewProcessDesignFile),
	protected(components.NewPage),
	protected(components.RemoveElementFromComponent),
	protected(components.CreateComponent),
	protected(jobs.NewPage),
	protected(templates.NewPage),
))
