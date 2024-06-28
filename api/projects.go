package api

import "github.com/labstack/echo/v4"

func NewProjectsController() ProjectsController {
	return ProjectsController{}
}

type ProjectsController struct{}

func (s ProjectsController) Load(e *echo.Echo) error {
	e.GET("/api/v1/projects", s.ListProjects())
	return nil
}

func (s ProjectsController) ListProjects() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
