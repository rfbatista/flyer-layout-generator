package projects

import (
	"algvisual/internal/ports"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		ports.AsController(NewProjectsController),
	),
)
