package projects

import (
	"algvisual/database"
	"algvisual/internal/infra/cognito"
	"algvisual/internal/infra/config"
	"algvisual/internal/infra/middlewares"
	"algvisual/internal/projects/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewProjectsController(
	db *database.Queries,
	cfg config.AppConfig,
	cog *cognito.Cognito,
) ProjectsController {
	return ProjectsController{
		db:  db,
		cfg: cfg,
		cog: cog,
	}
}

type ProjectsController struct {
	db  *database.Queries
	cfg config.AppConfig
	cog *cognito.Cognito
}

func (s ProjectsController) Load(e *echo.Echo) error {
	e.POST(
		"/api/v1/project",
		s.CreateProject(),
		middlewares.NewAuthMiddleware(s.cog, s.cfg),
	)
	e.GET(
		"/api/v1/projects",
		s.ListProjects(),
		middlewares.NewAuthMiddleware(s.cog, s.cfg),
	)
	e.GET(
		"/api/v1/project/:project_id",
		s.GetProjectByID(),
		middlewares.NewAuthMiddleware(s.cog, s.cfg),
	)
	e.PATCH(
		"/api/v1/project/:project_id",
		s.UpdateProject(),
		middlewares.NewAuthMiddleware(s.cog, s.cfg),
	)
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
