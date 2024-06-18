package worker

import (
	"context"

	"go.uber.org/fx"
)

var Module = fx.Options(fx.Provide(NewWorkerPool), fx.Invoke(RegisterHooks))

type RegisterHooksParams struct {
	fx.In
	WorkerPool WorkerPool
}

func RegisterHooks(lc fx.Lifecycle, params RegisterHooksParams) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			params.WorkerPool.Start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
