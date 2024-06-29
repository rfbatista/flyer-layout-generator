package api

import (
	"algvisual/internal/database"
	"algvisual/internal/templates"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func NewTemplatesController(
	db *database.Queries,
	pool *pgxpool.Pool,
	log *zap.Logger,
) TemplatesController {
	return TemplatesController{db: db, pool: pool, log: log}
}

type TemplatesController struct {
	db   *database.Queries
	pool *pgxpool.Pool
	log  *zap.Logger
}

func (s TemplatesController) Load(e *echo.Echo) error {
	e.GET("/api/v1/project/:project_id/templates", s.ListTemplates())
	e.POST("/api/v1/project/:project_id/templates", s.UploadTemplatesCSV())
	e.DELETE("/api/v1/project/:project_id/templates/:template_id", s.DeleteTemplate())
	return nil
}

func (s TemplatesController) ListTemplates() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req templates.ListTemplatesByProjectIdInput
		err := echo.PathParamsBinder(c).Int32("project_id", &req.ProjectID).BindError()
		if err != nil {
			return err
		}
		out, err := templates.ListTemplatesByProjectIdUseCase(c.Request().Context(), req, s.db)
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
		var req templates.TemplatesCsvUploadRequest
		err = c.Bind(&req)
		if err != nil {
			return err
		}
		req.File = &src
		out, err := templates.TemplatesCsvUploadUseCase(
			c.Request().Context(),
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
		var req templates.DeleteTemplateByIdInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := templates.DeleteTemplateByIdUseCase(c.Request().Context(), req, s.db)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}
