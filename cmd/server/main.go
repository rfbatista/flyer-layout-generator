package main

import (
	"algvisual/api"
	"algvisual/internal/advertisers"
	"algvisual/internal/clients"
	"algvisual/internal/infra"
	"algvisual/internal/layoutgenerator"
	"algvisual/internal/renderer"
	"algvisual/internal/templates"
	"algvisual/internal/worker"
	"fmt"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		api.Module,
		infra.Module,
		worker.Module,
		renderer.Module,
		templates.Module,
		layoutgenerator.Module,
		advertisers.Module,
		clients.Module,
		fx.NopLogger,
	)
	fmt.Println(app.Err())
	app.Run()
}
