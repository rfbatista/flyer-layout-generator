package infra

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupStaticServer(p HTTPServerParams, e *echo.Echo) {
	webStaticPath := fmt.Sprintf("%s/web/static", FindProjectRoot())
	webgroup := e.Group("/web")
	webgroup.Use(
		middleware.StaticWithConfig(middleware.StaticConfig{
			Root:   webStaticPath,
			Browse: true,
		}),
	)
	webDistPath := fmt.Sprintf("%s/dist/web", FindProjectRoot())
	distGroup := e.Group("/dist")
	distGroup.Use(
		middleware.StaticWithConfig(middleware.StaticConfig{
			Root:   webDistPath,
			Browse: true,
		}),
	)

	viteDistPath := fmt.Sprintf("%s/dist/vite", FindProjectRoot())
	viteDistGroup := e.Group("/dist/vite")
	viteDistGroup.Use(
		middleware.StaticWithConfig(middleware.StaticConfig{
			Root:   viteDistPath,
			Browse: true,
		}),
	)
}
