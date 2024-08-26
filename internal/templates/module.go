package templates

import (
	"algvisual/internal/ports"
	"algvisual/internal/templates/repository"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewTemplateService,
		ports.AsController(NewTemplatesController),
		repository.NewTemplateRepository,
	),
)
