package requestdefinecomponents

import (
	"algvisual/internal/designs"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/rfbatista/apitools"
	"go.uber.org/zap"

	"algvisual/internal/database"
	"algvisual/internal/shared"
	"algvisual/internal/web/components/notification"
)

func NewPage(db *database.Queries) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(shared.PageRequestElements.String())
	h.SetHandle(func(c echo.Context) error {
		sId := c.Param("id")
		if sId == "" {
			sId = "0"
		}
		id, err := strconv.ParseInt(sId, 10, 32)
		if err != nil {
			return shared.RenderComponent(
				shared.WithComponent(
					notification.FailureMessage(err.Error()), c,
				),
			)
		}
		out, err := PagePropsAssembler(c.Request().Context(), db, pageRequest{DesignID: int32(id)})
		if err != nil {
			return err
		}
		return shared.RenderComponent(
			shared.WithComponent(
				Page(int32(id), out),
				c,
			),
			shared.WithPage(shared.PageRequestUploadFile.String()),
		)
	})
	return h
}

func CreateComponent(db *database.Queries, tx *pgxpool.Pool, log *zap.Logger) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.POST)
	h.SetPath(shared.PageRequestElementsCreateComponent.String())
	h.SetHandle(func(c echo.Context) error {
		var req designs.CreateComponentRequest
		err := c.Bind(&req)
		if err != nil {
			return shared.RenderComponent(
				shared.WithComponent(
					notification.FailureMessage(err.Error()), c,
				),
			)
		}
		_, err = designs.CreateComponentUseCase(c.Request().Context(), req, db, tx, log)
		if err != nil {
			return err
		}
		c.Response().
			Header().
			Set("HX-Redirect", shared.PageRequestElements.Replace([]string{strconv.Itoa(int(req.DesignID))}))
		return c.NoContent(http.StatusOK)
	})
	return h
}

func RemoveElementFromComponent(db *database.Queries) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.POST)
	h.SetPath(shared.PageRequestElementsRemoveElement.String())
	h.SetHandle(func(c echo.Context) error {
		var req designs.RemoveComponentUseCaseRequest
		err := c.Bind(&req)
		if err != nil {
			return shared.RenderComponent(
				shared.WithComponent(
					notification.FailureMessage(err.Error()), c,
				),
			)
		}
		_, err = designs.RemoveComponentUseCase(c.Request().Context(), db, req)
		if err != nil {
			return err
		}
		c.Response().
			Header().
			Set("HX-Redirect", shared.PageRequestElements.Replace([]string{strconv.Itoa(int(req.DesignID))}))
		return c.NoContent(http.StatusOK)
	})
	return h
}
