package debug

import (
	"algvisual/internal/database"
	"algvisual/internal/entities"
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

func CreateImage(
	queries *database.Queries,
	conn *pgxpool.Pool,
	log *zap.Logger,
	db *pgxpool.Pool,
	config *infra.AppConfig,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.POST)
	h.SetPath("/editor/create/image/debug")
	h.SetHandle(func(c echo.Context) error {
		var req entities.Layout
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := layoutgenerator.GenerateImageFromLayoutUseCase(log, *config, layoutgenerator.GenerateImageFromLayoutInput{
			Layout:        req,
			DesignFileURL: "/home/renan/projetos/algvisual/banner/dist/files/Natura",
		})
		if err != nil {
			shared.ErrorNotification(c, err.Error())
			return c.NoContent(http.StatusBadRequest)
		}
		shared.SuccessNotification(c, "sucesso")
		return c.String(http.StatusOK, fmt.Sprintf(`
				<div class="center wrapper" style="height:450px;width:450px;">
					<img src="%s" class="img"/>
				</div>
		`, out.ImageURL))
	})
	return h
}
