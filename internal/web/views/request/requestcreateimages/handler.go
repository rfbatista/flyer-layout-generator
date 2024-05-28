package requestcreateimages

import (
	"algvisual/internal/database"
	"algvisual/internal/infra"
	"algvisual/internal/shared"
	"algvisual/internal/web/components/notification"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/rfbatista/apitools"
	"go.uber.org/zap"
)

func NewPage(
	db *database.Queries,
	client *infra.ImageGeneratorClient,
	log *zap.Logger,
	config *infra.AppConfig,
	pool *pgxpool.Pool,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(shared.PageRequestGenerateImages.String())
	h.SetHandle(func(c echo.Context) error {
		var req createImageRequest
		err := c.Bind(&req)
		if err != nil {
			return shared.RenderComponent(
				shared.WithComponent(
					notification.FailureMessage(err.Error()), c,
				),
			)
		}
		out, err := createImages(c.Request().Context(), req, db, client, log, *config, pool)
		if err != nil {
			return err
		}
		return shared.RenderComponent(
			shared.WithComponent(
				Page(*out),
				c,
			),
			shared.WithPage(shared.PageRequestUploadFile.String()),
		)
	})
	return h
}
