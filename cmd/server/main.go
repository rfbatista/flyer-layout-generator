package main

import (
	"algvisual/api"
	"algvisual/internal/advertisers"
	"algvisual/internal/clients"
	"algvisual/internal/designassets"
	"algvisual/internal/designprocessor"
	"algvisual/internal/designs"
	"algvisual/internal/iam"
	"algvisual/internal/infra"
	"algvisual/internal/layoutgenerator"
	"algvisual/internal/projects"
	"algvisual/internal/renderer"
	"algvisual/internal/templates"
	"algvisual/internal/worker"
	"fmt"
	"os"

	"go.uber.org/fx"
)

func main() {
	defer func() { // catch or finally
		if err := recover(); err != nil { // catch
			fmt.Fprintf(os.Stderr, "Exception: %v\n", err)
			os.Exit(1)
		}
	}()
	app := fx.New(
		api.Module,
		infra.Module,
		worker.Module,
		renderer.Module,
		templates.Module,
		layoutgenerator.Module,
		advertisers.Module,
		clients.Module,
		iam.Module,
		designprocessor.Module,
		projects.Module,
		designassets.Module,
		designs.Module,
	)
	fmt.Println(app.Err())
	app.Run()
}
