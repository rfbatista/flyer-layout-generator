package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rfbatista/apitools"
	"go.uber.org/zap"

	"algvisual/database"
	"algvisual/infra"
	"algvisual/shared"
	"algvisual/usecases"
)

func NewSavePhotoshopAPI(
	db *database.Queries,
	proc *infra.PhotoshopProcessor,
	storage infra.Storage,
	log *zap.Logger,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.POST)
	h.SetPath(shared.EndpointUploadPhotoshop.String())
	h.SetHandle(func(c echo.Context) error {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()
		req := usecases.SavePhotoshopFileUseCaseRequest{
			Filename: c.FormValue("filename"),
			File:     src,
		}
		out, err := usecases.SavePhotoshopFileUseCase(
			c.Request().Context(),
			db,
			req,
			storage,
			proc,
			log,
		)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	})
	return h
}

func NewListPhotoshop() apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(shared.EndpointListPhotoshop.String())
	h.SetHandle(func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	})
	return h
}
