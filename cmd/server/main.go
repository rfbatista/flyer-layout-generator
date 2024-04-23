package main

import (
	"algvisual/internal/api"
	"algvisual/internal/infra"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		api.Module, infra.Module,
	).Run()
}
