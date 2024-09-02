package controllers

import (
	"algvisual/internal/application/usecases/advertisers"
	"algvisual/internal/infrastructure/cognito"
	"algvisual/internal/infrastructure/config"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewAdvertiserController(
	db *database.Queries,
	ads advertisers.AdvertiserService,
	cfg config.AppConfig,
	cog *cognito.Cognito,
) AdvertiserController {
	return AdvertiserController{
		db:  db,
		ads: ads,
		cog: cog,
		cfg: cfg,
	}
}

type AdvertiserController struct {
	db  *database.Queries
	ads advertisers.AdvertiserService
	cfg config.AppConfig
	cog *cognito.Cognito
}

func (s AdvertiserController) Load(e *echo.Echo) error {
	e.GET(
		"/api/v1/advertisers",
		s.ListAdvertisers(),
		middlewares.NewAuthMiddleware(s.cog, s.cfg),
	)
	e.POST(
		"/api/v1/advertisers",
		s.CreateAdvertiser(),
		middlewares.NewAuthMiddleware(s.cog, s.cfg),
	)
	return nil
}

func (s AdvertiserController) ListAdvertisers() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req advertisers.ListAdvertisersInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := advertisers.ListAdvertisersUseCase(c, req, s.db)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s AdvertiserController) CreateAdvertiser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req advertisers.CreateAdvertiserInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := s.ads.Create(c, req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}
