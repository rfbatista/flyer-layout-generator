package clients

import (
	"algvisual/internal/application/usecases/advertisers"
	"context"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/gavv/httpexpect/v2"
	"go.uber.org/fx"
)

func TestCreateClient(tt *testing.T) {
	app := fx.New(
		advertisers.Module,
		fx.NopLogger,
	)
	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}
	e := httpexpect.Default(tt, "http://localhost:8000")
	tt.Run("should create a client", func(t *testing.T) {
		e.POST("/api/v1/advertisers").
			Expect().
			Status(http.StatusOK)
	})
}
