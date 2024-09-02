package main

import (
	"algvisual/internal/application/consumers"
	"algvisual/internal/application/controllers"
	"algvisual/internal/application/usecases/adaptations"
	"algvisual/internal/application/usecases/designassets"
	"algvisual/internal/application/usecases/designprocessor"
	"algvisual/internal/application/usecases/designs"
	"algvisual/internal/application/usecases/layoutgenerator"
	"algvisual/internal/application/usecases/projects"
	"algvisual/internal/application/usecases/renderer"
	"algvisual/internal/application/usecases/templates"
	"algvisual/internal/infrastructure"
	"algvisual/internal/infrastructure/repositories"
	"algvisual/internal/infrastructure/worker"
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
		consumers.Module,
		adaptations.Module,
		renderer.Module,
		layoutgenerator.Module,
		templates.Module,
		designprocessor.Module,
		designassets.Module,
		designs.Module,
		projects.Module,
		controllers.Module,
		infrastructure.Module,
		repositories.Module,
		worker.Module,
	)
	fmt.Println(app.Err())
	app.Run()
}
