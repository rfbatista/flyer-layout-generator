package infra

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Options(
	fx.Provide(
		NewLogger,
		NewConfig,
		NewHTTPServer,
		NewDatabaseConnection,
		NewDatabaseQueries,
		NewFileStorage,
		NewPhotoshpProcessor,
	),
	fx.Invoke(RegisterHooks),
)

type RegisterHooksParams struct {
	fx.In
	Server *echo.Echo
	Logger *zap.Logger
	Config *AppConfig
	Conn   *pgx.Conn
}

func RegisterHooks(lc fx.Lifecycle, params RegisterHooksParams) {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			params.Logger.Info(
				"starting http server",
				zap.String("port", params.Config.HTTPServer.Port),
			)
			go func() {
				err := params.Server.Start(fmt.Sprintf(":%s", params.Config.HTTPServer.Port))
				if err != nil {
					params.Logger.Info(
						"server startup failed", zap.Error(err),
					)
					return
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			err := params.Conn.Close(ctx)
			if err != nil {
				params.Logger.Error("failed to close database connection", zap.Error(err))
			}
			return params.Server.Shutdown(ctx)
		},
	})
}
