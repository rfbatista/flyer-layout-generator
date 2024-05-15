package requestdefinecomponents

import (
	"algvisual/internal/database"
	"algvisual/internal/shared"
	"algvisual/internal/web/components/notification"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rfbatista/apitools"
)

func NewPage(db *database.Queries) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(shared.PageRequestElements.String())
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
		d, err := db.GetdesignElements(c.Request().Context(), int32(id))
		if err != nil {
			return shared.RenderComponent(
				shared.WithComponent(
					notification.FailureMessage(err.Error()), c,
				),
			)
		}
		return shared.RenderComponent(
			shared.WithComponent(
				Page(d),
				c,
			),
			shared.WithPage(shared.PageRequestUploadFile.String()),
		)
	})
	return h
}
