package advertisers

import (
	"algvisual/internal/ports"

	"go.uber.org/fx"
)

func AsController(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(ports.Controller)),
		fx.ResultTags(`group:"controller"`),
	)
}

var Module = fx.Options(
	fx.Provide(
		NewAdvertiserService,
		NewAdvertiserRepository,
		AsController(NewAdvertiserController),
	),
)
