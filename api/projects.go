package api

import (
	"algvisual/internal/database"
	"algvisual/internal/projects"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewProjectsController(db *database.Queries) ProjectsController {
	return ProjectsController{db: db}
}

type ProjectsController struct {
	db *database.Queries
}

func (s ProjectsController) Load(e *echo.Echo) error {
	e.POST("/api/v1/project", s.CreateProject())
	e.GET("/api/v1/projects", s.ListProjects())
	e.GET("/api/v1/project/:project_id", s.GetProjectByID())
	e.PATCH("/api/v1/project/:project_id", s.UpdateProject())
	return nil
}

func (s ProjectsController) UpdateProject() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req projects.UpdateProjectByIdInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := projects.UpdateProjectByIdUseCase(c.Request().Context(), req, s.db)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s ProjectsController) CreateProject() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req projects.CreateProjectInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := projects.CreateProjectUseCase(c.Request().Context(), req, s.db)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s ProjectsController) ListProjects() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req projects.ListProjectsInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := projects.ListProjectsUseCase(c.Request().Context(), req, s.db)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s ProjectsController) GetProjectByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req projects.GetProjectByIdInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := projects.GetProjectByIdUseCase(c.Request().Context(), req, s.db)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}
