package controllers

import (
	"algvisual/internal/application/usecases/projects"
	"algvisual/internal/infrastructure/cognito"
	"algvisual/internal/infrastructure/config"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/middlewares"
	"algvisual/internal/shared"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewProjectsController(
	db *database.Queries,
	cfg config.AppConfig,
	cog *cognito.Cognito,
	listProjects projects.ListProjectsUseCase,
	getProjectById projects.GetProjectByIdUseCase,
	saveProjectLayout *projects.SaveProjectLayoutUseCase,
	listProjectLayouts *projects.ListProjectLayoutsUseCase,
) ProjectsController {
	return ProjectsController{
		db:                 db,
		cfg:                cfg,
		cog:                cog,
		listProjects:       listProjects,
		getProjectById:     getProjectById,
		saveProjectLayout:  saveProjectLayout,
		listProjectLayouts: listProjectLayouts,
	}
}

type ProjectsController struct {
	db                 *database.Queries
	cfg                config.AppConfig
	cog                *cognito.Cognito
	listProjectLayouts *projects.ListProjectLayoutsUseCase
	listProjects       projects.ListProjectsUseCase
	getProjectById     projects.GetProjectByIdUseCase
	saveProjectLayout  *projects.SaveProjectLayoutUseCase
}

func (s ProjectsController) Load(e *echo.Echo) error {
	// e.POST(
	// 	"/api/v1/project",
	// 	s.CreateProject(),
	// 	middlewares.NewAuthMiddleware(s.cog, s.cfg),
	// )
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
	e.POST(
		"/api/v1/project/:project_id/layout/:layout_id/save",
		s.SaveProjectLayouts(),
		middlewares.NewAuthMiddleware(s.cog, s.cfg),
	)
	e.GET(
		"/api/v1/project/:project_id/layouts",
		s.ListProjectLayouts(),
		middlewares.NewAuthMiddleware(s.cog, s.cfg),
	)
	// e.PATCH(
	// 	"/api/v1/project/:project_id",
	// 	s.UpdateProject(),
	// 	middlewares.NewAuthMiddleware(s.cog, s.cfg),
	// )
	// e.DELETE(
	// 	"/api/v1/project/:project_id",
	// 	s.DeleteProject(),
	// 	middlewares.NewAuthMiddleware(s.cog, s.cfg),
	// )
	return nil
}

// func (s ProjectsController) UpdateProject() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var req usecase.UpdateProjectByIdInput
// 		err := c.Bind(&req)
// 		if err != nil {
// 			return err
// 		}
// 		out, err := usecase.UpdateProjectByIdUseCase(c, req, s.db)
// 		if err != nil {
// 			return err
// 		}
// 		return c.JSON(http.StatusOK, out)
// 	}
// }

// func (s ProjectsController) CreateProject() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var req usecase.CreateProjectInput
// 		err := c.Bind(&req)
// 		if err != nil {
// 			return err
// 		}
// 		out, err := usecase.CreateProjectUseCase(c, req, s.db)
// 		if err != nil {
// 			return err
// 		}
// 		return c.JSON(http.StatusOK, out)
// 	}
// }

func (s ProjectsController) ListProjects() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req projects.ListProjectsInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		req.UserSession = *shared.GetSession(c)
		out, err := s.listProjects.Execute(c.Request().Context(), req)
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
		out, err := s.getProjectById.Execute(c.Request().Context(), req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

//
// func (s ProjectsController) DeleteProject() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var req usecase.DeleteProjectByIdInput
// 		err := c.Bind(&req)
// 		if err != nil {
// 			return err
// 		}
// 		out, err := usecase.DeleteProjectByIdUseCase(c, req, s.db)
// 		if err != nil {
// 			return err
// 		}
// 		return c.JSON(http.StatusOK, out)
// 	}
// }

func (s ProjectsController) SaveProjectLayouts() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req projects.SaveProjectLayoutInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := s.saveProjectLayout.Execute(c.Request().Context(), req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s ProjectsController) ListProjectLayouts() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req projects.ListProjectLayoutsInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := s.listProjectLayouts.Execute(c.Request().Context(), req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}
