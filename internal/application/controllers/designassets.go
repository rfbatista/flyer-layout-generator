package controllers

import (
	"algvisual/internal/application/usecases/designassets"
	"algvisual/internal/infrastructure/cognito"
	"algvisual/internal/infrastructure/config"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/middlewares"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewAssetsController(
	db *database.Queries,
	cfg config.AppConfig,
	s *designassets.DesignAssetService,
	cog *cognito.Cognito,
) AssetsController {
	return AssetsController{db: db, cfg: cfg, s: s, cog: cog}
}

type AssetsController struct {
	db  *database.Queries
	cfg config.AppConfig
	s   *designassets.DesignAssetService
	cog *cognito.Cognito
}

func (s AssetsController) Load(e *echo.Echo) error {
	e.POST(
		"/api/v1/images",
		s.UploadImage(),
	)
	e.GET(
		"/api/v1/project/:project_id/assets",
		s.GetProjectDesignAssets(),
		middlewares.NewAuthMiddleware(s.cog, s.cfg),
	)
	e.POST(
		"/api/v1/assets/:asset_id",
		s.AddAssetProperty(),
		middlewares.NewAuthMiddleware(s.cog, s.cfg),
	)
	e.GET(
		"/api/v1/images",
		s.ListGeneratedImages(),
		middlewares.NewAuthMiddleware(s.cog, s.cfg),
	)
	e.GET(
		"/api/v1/images/:image_name",
		s.DownloadImage(),
	)
	return nil
}

func (s AssetsController) UploadImage() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req designassets.ImageUploadRequest
		file, err := c.FormFile("file")
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		req.File = file
		req.Filename = c.FormValue("filename")
		out, err := s.s.ImageUpload(c, req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s AssetsController) GetProjectDesignAssets() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req designassets.GetDesignAssetByProjectIdInput
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
		var req designassets.AddNewAssetPropertyInput
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
		var req designassets.ListGeneratedImagesRequest
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
