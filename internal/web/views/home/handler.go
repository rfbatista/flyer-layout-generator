package home

import (
	"algvisual/internal/database"
	"algvisual/internal/layoutgenerator"
	"algvisual/internal/shared"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/rfbatista/apitools"
	"go.uber.org/zap"
)

func NewPageHome(
	queries *database.Queries,
	conn *pgxpool.Pool,
	log *zap.Logger,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(shared.PageHome.String())
	h.SetHandle(func(c echo.Context) error {
		out, err := layoutgenerator.ListLayout(c.Request().Context(), queries, 10, 0)
		if err != nil {
			log.Error("failed to render home page", zap.Error(err))
			return err
		}
		component := HomePage(out)
		w := c.Response().Writer
		err = component.Render(
			context.WithValue(c.Request().Context(), "page", shared.PageHome.String()),
			w,
		)
		if err != nil {
			log.Error("failed to render home page", zap.Error(err))
			return err
		}
		return nil
	})
	return h
}
