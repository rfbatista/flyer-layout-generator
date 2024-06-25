package debug

import (
	"algvisual/internal/database"
	"algvisual/internal/infra"
	"algvisual/internal/shared"
	"algvisual/web/render"
	"fmt"
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
	static, err := bundler.AddPage(infra.BundlerPageParams{
		EntryPoints: []string{
			fmt.Sprintf("%s/web/views/debug/index.js", infra.FindProjectRoot()),
		},
		Name: "debug",
	})
	if err != nil {
		panic(shared.WrapWithAppError(err, "failed to build web/view/debug page", ""))
	}
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath("/editor/:design_id/:layout_id/debug")
	h.SetHandle(func(c echo.Context) error {
		var req PageRequest
		err := c.Bind(&req)
		if err != nil {
			log.Error("failed to render debug page", zap.Error(err))
			return err
		}
		props, err := Props(c.Request().Context(), queries, req)
		if err != nil {
			log.Error("failed to render debug page props", zap.Error(err))
			return err
		}
		return render.Render(c, http.StatusOK, Page(props, static.CSSName, static.JSName))
	})
	return h
}
