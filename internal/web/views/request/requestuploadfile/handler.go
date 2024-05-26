package requestuploadfile

import (
	"algvisual/internal/designprocessor"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rfbatista/apitools"
	"go.uber.org/zap"

	"algvisual/internal/database"
	"algvisual/internal/infra"
	"algvisual/internal/shared"
	"algvisual/internal/web/components/notification"
)

func NewPageRequestUploadFile() apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(shared.PageRequestUploadFile.String())
	h.SetHandle(func(c echo.Context) error {
		return shared.RenderComponent(
			shared.WithComponent(
				PageRequestUploadFile(),
				c,
			),
			shared.WithPage(shared.PageRequestUploadFile.String()),
		)
	})
	return h
}

func NewUploadDesignAPI(
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
			return shared.RenderComponent(
				shared.WithComponent(
					notification.FailureMessage(err.Error()), c,
				),
			)
		}
		src, err := file.Open()
		if err != nil {
			return shared.RenderComponent(
				shared.WithComponent(
					notification.FailureMessage(err.Error()), c,
				),
			)
		}
		defer src.Close()
		req := designprocessor.UploadDesignFileUseCaseRequest{
			Filename: c.FormValue("filename"),
			File:     src,
		}
		out, err := designprocessor.UploadDesignFileUseCase(
			c.Request().Context(),
			db,
			req,
			storage.Upload,
			log,
		)
		if err != nil {
			return shared.RenderComponent(
				shared.WithComponent(
					notification.FailureMessage(err.Error()), c,
				),
			)
		}
		c.Response().
			Header().
			Set("HX-Redirect", shared.PageRequestProcessDesign.Replace([]string{strconv.Itoa(int(out.Design.ID))}))
		return nil
	})
	return h
}
