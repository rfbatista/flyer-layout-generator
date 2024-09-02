package renderer

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Options(
	fx.Provide(
		NewRendererService,
		NewTextDrawer,
	),
	fx.Invoke(RegisterHooks),
)

type RegisterHooksParams struct {
	fx.In
	TextDrawer *TextDrawer
	Logger     *zap.Logger
}

func RegisterHooks(lc fx.Lifecycle, params RegisterHooksParams) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			params.Logger.Info("loading fonts")
			return params.TextDrawer.LoadFonts()
		},
	})
}
