package designassets

import (
	"go.uber.org/fx"
)

var Module = fx.Options(fx.Provide(
	NewDesignAssetService,
))
