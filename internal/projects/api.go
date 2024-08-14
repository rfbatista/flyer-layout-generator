package projects

import (
	"algvisual/database"
	"algvisual/internal/projects/usecase"
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
		var req usecase.UpdateProjectByIdInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := usecase.UpdateProjectByIdUseCase(c, req, s.db)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s ProjectsController) CreateProject() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req usecase.CreateProjectInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := usecase.CreateProjectUseCase(c, req, s.db)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s ProjectsController) ListProjects() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req usecase.ListProjectsInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := usecase.ListProjectsUseCase(c, req, s.db)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s ProjectsController) GetProjectByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req usecase.GetProjectByIdInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := usecase.GetProjectByIdUseCase(c, req, s.db)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}
