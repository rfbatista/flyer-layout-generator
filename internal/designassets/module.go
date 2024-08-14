package designassets

import (
	"algvisual/internal/ports"

	"go.uber.org/fx"
)

var Module = fx.Options(fx.Provide(
	NewDesignAssetService,
	ports.AsController(NewAssetsController),
))
