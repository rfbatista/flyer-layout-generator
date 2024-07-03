package project

import (
	"algvisual/database"
	"algvisual/internal/infra"
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
	bundler *infra.Bundler,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath("/project")
	h.SetHandle(func(c echo.Context) error {
		return render.Render(c, http.StatusOK, Page())
	})
	return h
}
