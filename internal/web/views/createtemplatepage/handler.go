package createtemplatepage

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/rfbatista/apitools"
	"go.uber.org/zap"

	"algvisual/internal/database"
	"algvisual/internal/shared"
	"algvisual/web/views/proccessdesign"
)

func NewPageCreateTemplate(
	queries *database.Queries,
	conn *pgxpool.Pool,
	log *zap.Logger,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath("/criar")
	h.SetHandle(func(c echo.Context) error {
		component := CreateTemplatePage()
		w := c.Response().Writer
		err := component.Render(c.Request().Context(), w)
		if err != nil {
			log.Error("failed to render home page", zap.Error(err))
			return err
		}
		return nil
	})
	return h
}

func NewPageProccessDesign(
	queries *database.Queries,
	conn *pgxpool.Pool,
	log *zap.Logger,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(shared.PageProccessDesing.String())
	h.SetHandle(func(c echo.Context) error {
		ctx := c.Request().Context()
		d, err := queries.Listdesign(ctx, database.ListdesignParams{Offset: 0, Limit: 100})
		if err != nil {
			log.Error("failed to render home page", zap.Error(err))
			return err
		}
		component := proccessdesign.Page(d)
		w := c.Response().Writer
		err = component.Render(c.Request().Context(), w)
		if err != nil {
			log.Error("failed to render home page", zap.Error(err))
			return err
		}
		return nil
	})
	return h
}
