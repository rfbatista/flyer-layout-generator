package api

import (
	"algvisual/internal/database"
	"algvisual/internal/templates"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewTemplatesController(db *database.Queries) TemplatesController {
	return TemplatesController{db: db}
}

type TemplatesController struct {
	db *database.Queries
}

func (s TemplatesController) Load(e *echo.Echo) error {
	e.GET("/api/v1/project/:project_id/templates", s.ListTemplates())
	return nil
}

func (s TemplatesController) ListTemplates() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req templates.ListTemplatesByProjectIdInput
		err := c.Bind(&req)
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
