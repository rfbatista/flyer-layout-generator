package adaptations

import "go.uber.org/fx"

var Module = fx.Options(fx.Provide(
	NewGetActiveAdaptationBatchUseCase,
	NewStartAdaptationUseCase,
))
