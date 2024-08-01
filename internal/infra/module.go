package infra

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
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
		NewImageGenerator,
		NewServerSideEventManager,
		NewCognito,
	),
	fx.Invoke(RegisterHooks),
)

type RegisterHooksParams struct {
	fx.In
	Server  *echo.Echo
	Logger  *zap.Logger
	Config  *AppConfig
	Conn    *pgxpool.Pool
	Cognito *Cognito
	SSE     *ServerSideEventManager
}

func RegisterHooks(lc fx.Lifecycle, params RegisterHooksParams) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			params.Logger.Info("loading cognito jwk")
			err := params.Cognito.LoadJWK()
			if err != nil {
				return err
			}
			params.Logger.Info(
				"starting http server",
				zap.String("port", params.Config.HTTPServer.Port),
			)
			err = params.Conn.Ping(ctx)
			if err != nil {
				return err
			}
			go func() {
				err := params.Server.Start(fmt.Sprintf(":%s", params.Config.HTTPServer.Port))
				if err != nil {
					params.Logger.Info(
						"server startup failed", zap.Error(err),
					)
					return
				}
			}()
			go func() {
				params.SSE.Listen()
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			params.Logger.Info("closing server")
			params.Conn.Close()
			return params.Server.Shutdown(ctx)
		},
	})
}
