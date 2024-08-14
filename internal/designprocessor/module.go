package designprocessor

import (
	"algvisual/internal/ports"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewDesignProcessorService,
		ports.AsController(NewDesignController),
	),
)
