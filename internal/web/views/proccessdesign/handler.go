package proccessdesign

import (
	"github.com/labstack/echo/v4"
	"github.com/rfbatista/apitools"
	"go.uber.org/zap"

	"algvisual/internal/database"
	"algvisual/internal/infra"
	"algvisual/internal/shared"
	"algvisual/internal/usecases"
	"algvisual/internal/web/components/notification"
)

func NewWebUploadDesignAPI(
	db *database.Queries,
	proc *infra.PhotoshopProcessor,
	storage infra.FileStorage,
	log *zap.Logger,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.POST)
	h.SetPath(shared.WebEndpointUploadPhotoshop.String())
	h.SetHandle(func(c echo.Context) error {
		file, err := c.FormFile("file")
		if err != nil {
			return shared.RenderComponent(notification.FailureMessage(err.Error()), c)
		}
		src, err := file.Open()
		if err != nil {
			return shared.RenderComponent(notification.FailureMessage(err.Error()), c)
		}
		defer src.Close()
		req := usecases.UploadDesignFileUseCaseRequest{
			Filename: c.FormValue("filename"),
			File:     src,
		}
		_, err = usecases.UploadDesignFileUseCase(
			c.Request().Context(),
			db,
			req,
			storage.Upload,
			log,
		)
		if err != nil {
			return shared.RenderComponent(notification.FailureMessage(err.Error()), c)
		}
		return shared.RenderComponent(
			notification.SuccessMessage("Arquivo cadastrado com sucesso"),
			c,
		)
	})
	return h
}
