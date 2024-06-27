package main

import (
	"algvisual/api"
	"algvisual/internal/infra"
	"algvisual/internal/worker"
	"algvisual/web"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		api.Module, infra.Module, web.Module, worker.Module,
	).Run()
}
