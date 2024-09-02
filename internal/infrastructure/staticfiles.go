package infrastructure

import (
	"algvisual/internal/infrastructure/config"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupStaticServer(p HTTPServerParams, e *echo.Echo) {
	webStaticPath := fmt.Sprintf("%s/web/static", config.FindProjectRoot())
	webgroup := e.Group("/web")
	webgroup.Use(
		middleware.StaticWithConfig(middleware.StaticConfig{
			Root:   webStaticPath,
			Browse: true,
		}),
	)
	webDistPath := fmt.Sprintf("%s/dist/web", config.FindProjectRoot())
	distGroup := e.Group("/dist")
	distGroup.Use(
		middleware.StaticWithConfig(middleware.StaticConfig{
			Root:   webDistPath,
			Browse: true,
		}),
	)

	viteDistPath := fmt.Sprintf("%s/dist/vite", config.FindProjectRoot())
	viteDistGroup := e.Group("/dist/vite")
	viteDistGroup.Use(
		middleware.StaticWithConfig(middleware.StaticConfig{
			Root:   viteDistPath,
			Browse: true,
		}),
	)
}
