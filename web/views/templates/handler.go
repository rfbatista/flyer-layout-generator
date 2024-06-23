package templates

import (
	"algvisual/internal/database"
	"algvisual/internal/infra"
	"algvisual/internal/shared"
	"algvisual/internal/templates"
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
			fmt.Sprintf("%s/web/views/templates/index.js", infra.FindProjectRoot()),
		},
		Name: "files/editor",
	})
	if err != nil {
		panic(shared.WrapWithAppError(err, "failed to build editor new doc page", ""))
	}
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(shared.PageListTemplate.String())
	h.SetHandle(func(c echo.Context) error {
		var req pageProps
		err := c.Bind(&req)
		if err != nil {
			log.Error("failed to render home page", zap.Error(err))
			return err
		}
		props, err := Props(c.Request().Context(), queries, log, req)
		if err != nil {
			log.Error("failed to render home page", zap.Error(err))
			return err
		}
		return render.Render(c, http.StatusOK, Page(props, static.CSSName, static.JSName))
	})
	return h
}

func CreateTemplate(
	queries *database.Queries,
	conn *pgxpool.Pool,
	log *zap.Logger,
	bundler *infra.Bundler,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.POST)
	h.SetPath("/template")
	h.SetHandle(func(c echo.Context) error {
		var req templates.CreateTemplateUseCaseRequest
		err := c.Bind(&req)
		if err != nil {
			log.Error("failed to render home page", zap.Error(err))
			return err
		}
		_, err = templates.CreateTemplateUseCase(c.Request().Context(), conn, queries, req, log)
		if err != nil {
			log.Error("failed to create template", zap.Error(err))
			render.ErrorNotification(c, err.Error())
			return c.NoContent(http.StatusBadRequest)
		}
		render.SuccessNotification(c, "Template criado com sucesso")
		return c.NoContent(http.StatusOK)
	})
	return h
}
