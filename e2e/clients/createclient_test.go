package clients

import (
	"algvisual/internal/advertisers"
	"algvisual/internal/infra"
	"context"
	"log"
	"net/http"
	"testing"
	"time"

	"go.uber.org/fx"
)

func TestCreateClient(tt *testing.T) {
	app := fx.New(
		infra.TestModule,
		advertisers.Module,
		fx.NopLogger,
	)
	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}
	tt.Run("should create a client", func(t *testing.T) {
		if _, err := http.Get("http://localhost:8000/"); err != nil {
			log.Fatal(err)
		}
		stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		if err := app.Stop(stopCtx); err != nil {
			log.Fatal(err)
		}
	})
}
