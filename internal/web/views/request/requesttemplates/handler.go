package requesttemplates

import (
	"algvisual/internal/templates"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/rfbatista/apitools"
	"go.uber.org/zap"

	"algvisual/internal/database"
	"algvisual/internal/shared"
	"algvisual/internal/web/components/notification"
)

func NewPage(db *database.Queries) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(string(shared.PageRequestUploadSheet.String()))
	h.SetHandle(func(c echo.Context) error {
		sId := c.Param("design_id")
		return shared.RenderComponent(
			shared.WithComponent(
				Page(sId),
				c,
			),
			shared.WithPage(shared.PageRequestUploadFile.String()),
		)
	})
	return h
}

func NewPageTemplatesCreated(db *database.Queries) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(shared.PageRequestTemplatesCreated.String())
	h.SetHandle(func(c echo.Context) error {
		reqID := c.Param("request_id")
		designID := c.Param("design_id")
		templates, err := db.GetTemplatesByRequestID(
			c.Request().Context(),
			pgtype.Text{String: reqID, Valid: true},
		)
		if err != nil {
			return shared.RenderComponent(
				shared.WithComponent(
					notification.FailureMessage(err.Error()), c,
				),
			)
		}
		return shared.RenderComponent(
			shared.WithComponent(
				TemplatesCreated(reqID, designID, templates),
				c,
			),
			shared.WithPage(shared.PageRequestUploadFile.String()),
		)
	})
	return h
}

func NewUploadCSV(
	db *database.Queries,
	pool *pgxpool.Pool,
	log *zap.Logger,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.POST)
	h.SetPath(shared.PageRequestUploadSheetCreateTemplates.String())
	h.SetHandle(func(c echo.Context) error {
		sId := c.Param("design_id")
		file, err := c.FormFile("file")
		if err != nil {
			return shared.RenderComponent(
				shared.WithComponent(
					notification.FailureMessage(err.Error()), c,
				),
			)
		}
		src, err := file.Open()
		if err != nil {
			return shared.RenderComponent(
				shared.WithComponent(
					notification.FailureMessage(err.Error()), c,
				),
			)
		}
		out, err := templates.TemplatesCsvUploadUseCase(
			c.Request().Context(),
			templates.TemplatesCsvUploadRequest{
				File: &src,
			},
			pool,
			db,
			log,
		)
		c.Response().
			Header().
			Set("HX-Redirect", shared.PageRequestTemplatesCreated.Replace([]string{sId, out.RequestID}))
		return nil
	})
	return h
}
