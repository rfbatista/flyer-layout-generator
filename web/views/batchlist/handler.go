package batchlist

import (
	"algvisual/internal/database"
	"algvisual/web/render"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/rfbatista/apitools"
	"go.uber.org/zap"
)

func NewPage(
	queries *database.Queries,
	conn *pgxpool.Pool,
	log *zap.Logger,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath("/lotes")
	h.SetHandle(func(c echo.Context) error {
		props, err := Props(c, queries, log)
		if err != nil {
			log.Error("failed to render lotes page", zap.Error(err))
			return err
		}
		return render.Render(c, http.StatusOK, Page(props))
	})
	return h
}

func NewTable(
	queries *database.Queries,
	conn *pgxpool.Pool,
	log *zap.Logger,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath("/lotes/table")
	h.SetHandle(func(c echo.Context) error {
		props, err := Props(c, queries, log)
		if err != nil {
			log.Error("failed to render lotes page", zap.Error(err))
			return err
		}
		return render.Render(c, http.StatusOK, Table(props))
	})
	return h
}
