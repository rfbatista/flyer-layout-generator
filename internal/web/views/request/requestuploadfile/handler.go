package requestuploadfile

import (
	"github.com/labstack/echo/v4"
	"github.com/rfbatista/apitools"

	"algvisual/internal/shared"
)

func NewPageRequestUploadFile() apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(shared.PageRequestUploadFile.String())
	h.SetHandle(func(c echo.Context) error {
		return shared.RenderComponent(
			PageRequestUploadFile(),
			c,
		)
	})
	return h
}
