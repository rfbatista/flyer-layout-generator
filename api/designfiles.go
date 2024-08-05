package api

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rfbatista/apitools"

	"algvisual/database"
	"algvisual/internal/infra/config"
	"algvisual/internal/shared"
)

func NewDownloadDesignFiles(
	cfg *config.AppConfig,
	db *database.Queries,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(shared.DownloadDesignFileEndpoint.String())
	h.SetHandle(func(c echo.Context) error {
		id := c.Param("id")
		idInt, err := strconv.ParseInt(id, 10, 32)
		if err != nil {
			return err
		}
		design, err := db.Getdesign(c.Request().Context(), int32(idInt))
		if err != nil {
			return err
		}
		return c.File(design.FileUrl.String)
	})
	return h
}
