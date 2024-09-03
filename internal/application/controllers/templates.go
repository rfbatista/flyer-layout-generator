package controllers

import (
	"algvisual/internal/application/usecases/templates"
	usecase "algvisual/internal/application/usecases/templates"
	"algvisual/internal/infrastructure/cognito"
	"algvisual/internal/infrastructure/config"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/middlewares"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func NewTemplatesController(
	db *database.Queries,
	pool *pgxpool.Pool,
	log *zap.Logger,
	cfg config.AppConfig,
	cog *cognito.Cognito,
) TemplatesController {
	return TemplatesController{db: db, pool: pool, log: log, cog: cog, cfg: cfg}
}

type TemplatesController struct {
	db   *database.Queries
	pool *pgxpool.Pool
	log  *zap.Logger
	cfg  config.AppConfig
	cog  *cognito.Cognito
}

func (s TemplatesController) Load(e *echo.Echo) error {
	e.GET(
		"/api/v1/project/:project_id/templates",
		s.ListTemplates(),
		middlewares.NewAuthMiddleware(s.cog, s.cfg),
	)
	e.GET(
		"/api/v1/template/:template_id",
		s.GetTemplateByID(),
		middlewares.NewAuthMiddleware(s.cog, s.cfg),
	)
	e.POST(
		"/api/v1/project/:project_id/templates",
		s.UploadTemplatesCSV(),
		middlewares.NewAuthMiddleware(s.cog, s.cfg),
	)
	e.DELETE(
		"/api/v1/project/:project_id/templates/:template_id",
		s.DeleteTemplate(),
		middlewares.NewAuthMiddleware(s.cog, s.cfg),
	)
	return nil
}

func (s TemplatesController) ListTemplates() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req usecase.ListTemplatesByProjectIdInput
		err := echo.PathParamsBinder(c).Int32("project_id", &req.ProjectID).BindError()
		if err != nil {
			return err
		}
		out, err := usecase.ListTemplatesByProjectIdUseCase(c, req, s.db)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s TemplatesController) UploadTemplatesCSV() echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()
		var req usecase.TemplatesCsvUploadRequest
		err = c.Bind(&req)
		if err != nil {
			return err
		}
		req.File = &src
		out, err := usecase.TemplatesCsvUploadUseCase(
			c,
			req,
			s.pool,
			s.db,
			s.log,
		)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s TemplatesController) DeleteTemplate() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req usecase.DeleteTemplateByIdInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := usecase.DeleteTemplateByIdUseCase(c, req, s.db)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s TemplatesController) GetTemplateByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req templates.GetTemplateByIdInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := templates.GetTemplateByIdUseCase(c.Request().Context(), req, s.db)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}
