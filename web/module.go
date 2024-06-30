package web

import (
	"algvisual/internal/ports"
	"algvisual/web/pages/editor"
	"algvisual/web/pages/generation"
	"algvisual/web/pages/project"
	"algvisual/web/pages/projects"
	"algvisual/web/views/batchlist"
	"algvisual/web/views/batchresults"
	"algvisual/web/views/components"
	"algvisual/web/views/debug"
	"algvisual/web/views/files"
	"algvisual/web/views/generate"
	"algvisual/web/views/jobs"
	"algvisual/web/views/modifier"
	"algvisual/web/views/single"
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
	protected(single.NewPageHome),
	protected(single.CreateImage),
	protected(files.NewPage),
	protected(files.NewUploadDesignAPI),
	protected(files.NewProcessDesignFile),
	protected(components.NewPage),
	protected(components.RemoveElementFromComponent),
	protected(components.CreateComponent),
	protected(jobs.NewPage),
	protected(templates.NewPage),
	protected(templates.CreateTemplate),
	protected(generate.NewPage),
	protected(generate.CreateRequest),
	protected(generate.CreateImage),
	protected(generate.CreateComponent),
	protected(batchlist.NewPage),
	protected(batchlist.NewTable),
	protected(batchresults.NewPage),
	protected(modifier.NewPage),
	protected(debug.NewPage),
	protected(debug.CreateImage),
	protected(projects.NewPage),
	protected(project.NewPage),
	protected(editor.NewPage),
	protected(generation.NewPage),
))
