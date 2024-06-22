package modifier

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
	pageentry := fmt.Sprintf("%s/web/views/modifier/index.js", infra.FindProjectRoot())
	static, err := bundler.AddPage(infra.BundlerPageParams{
		EntryPoints: []string{pageentry},
		Name:        "editor/newdoc",
	})
	if err != nil {
		panic(shared.WrapWithAppError(err, "failed to build editor new doc page", ""))
	}
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath("/batch/result/:layout_id")
	h.SetHandle(func(c echo.Context) error {
		var req request
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		props, err := Props(c.Request().Context(), queries, log, req)
		if err != nil {
			log.Error("failed to render home page", zap.Error(err))
			return err
		}
		fmt.Println(static.JSPath)
		return render.Render(
			c,
			http.StatusOK,
			Page(props, []string{static.JSName}, []string{static.CSSName}),
		)
	})
	return h
}
