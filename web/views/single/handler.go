package single

import (
	"algvisual/database"
	"algvisual/internal/infra"
	"algvisual/internal/layoutgenerator"
	"algvisual/internal/renderer"
	"algvisual/internal/shared"
	"algvisual/web/render"
	"net/http"

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
	h.SetPath(shared.PageSingleJob.String())
	h.SetHandle(func(c echo.Context) error {
		props, err := Props(c.Request().Context(), queries, log)
		if err != nil {
			log.Error("failed to render home page", zap.Error(err))
			return err
		}
		return render.Render(c, http.StatusOK, HomePage(props))
	})
	return h
}

func CreateImage(
	queries *database.Queries,
	conn *pgxpool.Pool,
	log *zap.Logger,
	db *pgxpool.Pool,
	config *infra.AppConfig,
	r renderer.RendererService,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.POST)
	h.SetPath(shared.PageHomeCreateImage.String())
	h.SetHandle(func(c echo.Context) error {
		var req layoutgenerator.GenerateImage
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := layoutgenerator.GenerateImageUseCase(
			c.Request().Context(),
			req,
			queries,
			db,
			*config,
			log,
			r,
		)
		if err != nil {
			shared.ErrorNotification(c, err.Error())
			return c.NoContent(http.StatusBadRequest)
		}
		shared.SuccessNotification(c, "sucesso")
		return shared.RenderComponent(
			shared.WithComponent(Image(out.Data.ImageURL, 100, 100), c),
		)
	})
	return h
}
