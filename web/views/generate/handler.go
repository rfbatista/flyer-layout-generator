package generate

import (
	"algvisual/internal/database"
	"algvisual/internal/infra"
	"algvisual/internal/layoutgenerator"
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
			fmt.Sprintf("%s/web/views/generate/index.js", infra.FindProjectRoot()),
		},
		Name: "editor/newdoc",
	})
	if err != nil {
		panic(shared.WrapWithAppError(err, "failed to build editor new doc page", ""))
	}
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath("/editor")
	h.SetHandle(func(c echo.Context) error {
		props, err := Props(c.Request().Context(), queries, log)
		if err != nil {
			log.Error("failed to render home page", zap.Error(err))
			return err
		}
		return render.Render(c, http.StatusOK, Page(props, static.CSSName, static.JSName))
	})
	return h
}

func CreateRequest(
	queries *database.Queries,
	conn *pgxpool.Pool,
	log *zap.Logger,
	db *pgxpool.Pool,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.POST)
	h.SetPath("/request/batch")
	h.SetHandle(func(c echo.Context) error {
		var req layoutgenerator.CreateLayoutRequestInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := layoutgenerator.CreateLayoutRequestUseCase(
			c.Request().Context(),
			queries,
			db,
			req,
		)
		if err != nil {
			shared.ErrorNotification(c, err.Error())
			return c.NoContent(http.StatusBadRequest)
		}
		shared.SuccessNotification(c, "sucesso")
		return c.JSON(http.StatusOK, out)
	})
	return h
}
