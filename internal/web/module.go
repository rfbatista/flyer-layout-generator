package web

import (
	"algvisual/internal/ports"
	"algvisual/internal/web/views/components"
	"algvisual/internal/web/views/files"
	"algvisual/internal/web/views/home"
	"algvisual/internal/web/views/jobs"
	"algvisual/internal/web/views/request/requestcreateimages"
	"algvisual/internal/web/views/request/requestdefinecomponents"
	"algvisual/internal/web/views/request/requestprocessdesign"
	"algvisual/internal/web/views/request/requesttemplates"
	"algvisual/internal/web/views/request/requestuploadfile"
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
	protected(requestuploadfile.NewPageRequestUploadFile),
	protected(requestuploadfile.NewUploadDesignAPI),
	protected(requestprocessdesign.NewPageRequestProcessDesign),
	protected(requestprocessdesign.NewWebProccessDesign),
	protected(requestdefinecomponents.NewPage),
	protected(requestdefinecomponents.CreateComponent),
	protected(requestdefinecomponents.RemoveElementFromComponent),
	protected(requesttemplates.NewPage),
	protected(requesttemplates.NewPageTemplatesCreated),
	protected(requesttemplates.NewUploadCSV),
	protected(requestcreateimages.NewPage),
	protected(files.NewPage),
	protected(home.CreateRequest),
	protected(components.NewPage),
	protected(jobs.NewPage),
))
