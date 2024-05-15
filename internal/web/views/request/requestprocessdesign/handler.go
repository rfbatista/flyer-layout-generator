package requestprocessdesign

import (
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/rfbatista/apitools"
	"go.uber.org/zap"

	"algvisual/internal/database"
	"algvisual/internal/infra"
	"algvisual/internal/shared"
	"algvisual/internal/usecases"
	"algvisual/internal/web/components/notification"
)

func NewPageRequestProcessDesign(db *database.Queries) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(shared.PageRequestProcessDesign.String())
	h.SetHandle(func(c echo.Context) error {
		sId := c.Param("id")
		id, err := strconv.ParseInt(sId, 10, 32)
		if err != nil {
			return shared.RenderComponent(
				shared.WithComponent(
					notification.FailureMessage(err.Error()), c,
				),
			)
		}
		d, err := db.Getdesign(c.Request().Context(), int32(id))
		if err != nil {
			return shared.RenderComponent(
				shared.WithComponent(
					notification.FailureMessage(err.Error()), c,
				),
			)
		}
		return shared.RenderComponent(
			shared.WithComponent(
				Page([]database.Design{d}),
				c,
			),
			shared.WithPage(shared.PageRequestUploadFile.String()),
		)
	})
	return h
}

func NewWebProccessDesign(
	db *database.Queries,
	proc *infra.PhotoshopProcessor,
	storage infra.FileStorage,
	log *zap.Logger,
	pool *pgxpool.Pool,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.POST)
	h.SetPath(shared.WebEndpointProccessDesign.String())
	h.SetHandle(func(c echo.Context) error {
		var req usecases.ProcessDesignFileRequest
		err := c.Bind(&req)
		if err != nil {
			c.Response().Header().Set("HX-Trigger", shared.InfoNotificationMessage(err.Error()))
			return c.NoContent(http.StatusOK)
		}
		_, err = usecases.ProcessDesignFileUseCase(
			c.Request().Context(),
			req,
			proc.ProcessFile,
			log,
			db,
			pool,
		)
		if err != nil {
			c.Response().Header().Set("HX-Trigger", shared.InfoNotificationMessage(err.Error()))
			return c.NoContent(http.StatusOK)
		}
		c.Response().Header().Set("HX-Trigger", shared.InfoNotificationMessage("Processo realizado com sucesso"))
		c.Response().
			Header().
			Set("HX-Redirect", shared.PageRequestElements.Replace([]string{strconv.Itoa(int(req.ID))}))
		return c.NoContent(http.StatusOK)
	})
	return h
}
