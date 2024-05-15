package main

import (
	"go.uber.org/fx"

	"algvisual/internal/api"
	"algvisual/internal/infra"
	"algvisual/internal/web"
)

func main() {
	fx.New(
		api.Module, infra.Module, web.Module,
	).Run()
}
