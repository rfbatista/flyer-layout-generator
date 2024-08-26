package layoutgenerator

import (
	"algvisual/internal/layoutgenerator/repository"
	"algvisual/internal/ports"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		ports.AsController(NewLayoutController),
		NewLayoutGeneratorService,
		repository.NewLayoutJobRepository,
		repository.NewRepository,
	),
)
