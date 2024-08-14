package designassets

import (
	"algvisual/database"
	"algvisual/internal/designassets/usecase"
	"algvisual/internal/infra/config"
	"algvisual/internal/shared"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewAssetsController(
	db *database.Queries,
	cfg *config.AppConfig,
	s *DesignAssetService,
) AssetsController {
	return AssetsController{db: db, cfg: cfg, s: s}
}

type AssetsController struct {
	db  *database.Queries
	cfg *config.AppConfig
	s   *DesignAssetService
}

func (s AssetsController) Load(e *echo.Echo) error {
	e.GET("/api/v1/images", s.UploadImage())
	e.POST("/api/v1/images", s.UploadImage())
	e.GET("/api/v1/project/:project_id/assets", s.GetProjectDesignAssets())
	e.POST("/api/v1/assets/:asset_id", s.AddAssetProperty())
	e.GET(shared.EndpointListImagesGenerated.String(), s.ListGeneratedImages())
	e.GET(shared.DownloadImageEndpoint.String(), s.DownloadImage())
	return nil
}

func (s AssetsController) UploadImage() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req usecase.ImageUploadRequest
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := s.s.ImageUpload(c, req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s AssetsController) GetProjectDesignAssets() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req usecase.GetDesignAssetByProjectIdInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := s.s.GetDesignAssetByProjectID(
			c,
			req,
		)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s AssetsController) AddAssetProperty() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req usecase.AddNewAssetPropertyInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := s.s.AddNewAssetProperty(
			c,
			req,
		)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s AssetsController) ListGeneratedImages() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req usecase.ListGeneratedImagesRequest
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := s.s.ListGeneratedImages(
			c,
			req,
		)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s AssetsController) DownloadImage() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("image_name")
		return c.File(fmt.Sprintf("%s/%s", s.cfg.ImagesFolderPath, name))
	}
}
