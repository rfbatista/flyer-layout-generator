package controllers

import (
	"algvisual/internal/ports"

	"go.uber.org/fx"
)

func AsController(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(ports.Controller)),
		fx.ResultTags(`group:"controller"`),
	)
}

var Module = fx.Options(
	fx.Provide(
		AsController(NewProjectsController),
		AsController(NewDesignController),
		AsController(NewAssetsController),
		AsController(NewDesignProcessorController),
		AsController(NewTemplatesController),
		AsController(NewLayoutController),
		AsController(NewJobsController),
		AsController(NewClientsController),
		AsController(NewAdvertiserController),
		AsController(NewIAMController),
	),
)
