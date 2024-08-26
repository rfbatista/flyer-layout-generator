package main

import (
	"algvisual/api"
	"algvisual/internal/adaptations"
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
	"log"
	"os"
	"runtime"

	"go.uber.org/fx"
)

func main() {
	defer func() { // catch or finally
		if err := recover(); err != nil { // catch
			fmt.Fprintf(os.Stderr, "exception: %v\n", err)
			buf := make([]byte, 1024)
			n := runtime.Stack(buf, false)
			log.Printf("Stack trace:\n%s", buf[:n])
			os.Exit(1)
		}
	}()
	app := fx.New(
		worker.Module,
		adaptations.Module,
		api.Module,
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
		infra.Module,
	)
	fmt.Println(app.Err())
	app.Run()
}
