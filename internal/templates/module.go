package templates

import (
	"algvisual/internal/ports"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewTemplateService,
		ports.AsController(NewTemplatesController),
	),
)
