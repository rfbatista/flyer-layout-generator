package designs

import (
	"algvisual/internal/ports"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewDesignService,
		ports.AsController(NewDesignController),
	),
)
