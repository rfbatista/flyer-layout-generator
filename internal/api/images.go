package api

import (
	"algvisual/internal/designassets"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/rfbatista/apitools"
	"go.uber.org/zap"

	"algvisual/internal/database"
	"algvisual/internal/infra"
	"algvisual/internal/shared"
)

func NewListGeneratedImagesAPI(
	queries *database.Queries,
	conn *pgxpool.Pool,
	log *zap.Logger,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(shared.EndpointListImagesGenerated.String())
	h.SetHandle(func(c echo.Context) error {
		var req designassets.ListGeneratedImagesRequest
		err := c.Bind(&req)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		result, err := designassets.ListGeneratedImagesUseCase(c.Request().Context(), req, queries, log)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	})
	return h
}

func NewUploadImage(
	cfg *infra.AppConfig,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.POST)
	h.SetPath(shared.UploadImageEndpoint.String())
	h.SetHandle(func(c echo.Context) error {
		var req designassets.ImageUploadRequest
		file, err := c.FormFile("file")
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		req.File = file
		req.Filename = c.FormValue("filename")
		result, err := designassets.ImageUploadUseCase(c.Request().Context(), req, cfg)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	})
	return h
}

func NewDownloadImage(
	cfg *infra.AppConfig,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(shared.DownloadImageEndpoint.String())
	h.SetHandle(func(c echo.Context) error {
		name := c.Param("image_name")
		return c.File(fmt.Sprintf("%s/%s", cfg.ImagesFolderPath, name))
	})
	return h
}
