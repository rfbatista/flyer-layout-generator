package main

import (
	"algvisual/api"
	"algvisual/internal/infra"
	"algvisual/internal/layoutgenerator"
	"algvisual/internal/renderer"
	"algvisual/internal/templates"
	"algvisual/internal/worker"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		api.Module,
		infra.Module,
		worker.Module,
		renderer.Module,
		templates.Module,
		layoutgenerator.Module,
		fx.NopLogger,
	).Run()
}
