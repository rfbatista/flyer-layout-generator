package main

import (
	"algvisual/internal/api"
	"algvisual/internal/infra"
	"algvisual/internal/web"
	"algvisual/internal/worker"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		api.Module, infra.Module, web.Module, worker.Module,
	).Run()
}
