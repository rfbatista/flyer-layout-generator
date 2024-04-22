package main

import (
	"go.uber.org/fx"

	"algvisual/api"
	"algvisual/infra"
)

func main() {
	fx.New(
		api.Module, infra.Module,
	).Run()
}
