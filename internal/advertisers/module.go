package advertisers

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewAdvertiserService,
		NewAdvertiserRepository,
	),
)
