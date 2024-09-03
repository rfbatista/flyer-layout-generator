package replications

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewStartReplicationUseCase,
		NewGetActiveReplicationUseCase,
	),
)
