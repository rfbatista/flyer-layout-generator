package api

import (
	"algvisual/database"
	"algvisual/internal/designassets"
	"algvisual/internal/infra"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewAssetsController(db *database.Queries, cfg *infra.AppConfig) AssetsController {
	return AssetsController{db: db, cfg: cfg}
}

type AssetsController struct {
	db  *database.Queries
	cfg *infra.AppConfig
}

func (s AssetsController) Load(e *echo.Echo) error {
	e.GET("/api/v1/images", s.UploadImage())
	return nil
}

func (s AssetsController) UploadImage() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req designassets.ImageUploadRequest
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := designassets.ImageUploadUseCase(c.Request().Context(), req, s.cfg)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}
