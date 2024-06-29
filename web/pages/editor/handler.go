package editor

import (
	"algvisual/internal/database"
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
	h.SetPath("/editor")
	h.SetHandle(func(c echo.Context) error {
		var req request
		err := c.Bind(&req)
		if err != nil {
			log.Error("failed to render editor page", zap.Error(err))
			return err
		}
		props, err := Props(c.Request().Context(), queries, log, req)
		if err != nil {
			log.Error("failed to render editor page props", zap.Error(err))
			return err
		}
		return render.Render(c, http.StatusOK, Page(props))
	})
	return h
}
