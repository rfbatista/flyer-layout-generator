package components

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
	"algvisual/web/components/notification"
)

func NewPage(db *database.Queries) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(shared.PageDefineComponents.String())
	h.SetHandle(func(c echo.Context) error {
		var req pageRequest
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := Props(c.Request().Context(), db, req)
		if err != nil {
			return err
		}
		return shared.RenderComponent(
			shared.WithComponent(
				Page(req.DesignID, out),
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
	h.SetPath("/design/:design_id/layout/:layout_id/component")
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
			Set("HX-Redirect", shared.PageDefineComponents.Replace([]string{strconv.Itoa(int(req.DesignID))}))
		return c.NoContent(http.StatusOK)
	})
	return h
}

func RemoveElementFromComponent(db *database.Queries) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.POST)
	h.SetPath(shared.PageDefineComponentsRemove.String())
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
			Set("HX-Redirect", shared.PageDefineComponents.Replace([]string{strconv.Itoa(int(req.DesignID))}))
		return c.NoContent(http.StatusOK)
	})
	return h
}
