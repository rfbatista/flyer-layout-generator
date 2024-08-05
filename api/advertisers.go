package api

import (
	"algvisual/database"
	"algvisual/internal/advertisers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewAdvertiserController(
	db *database.Queries,
	ads advertisers.AdvertiserService,
) AdvertiserController {
	return AdvertiserController{db: db, ads: ads}
}

type AdvertiserController struct {
	db  *database.Queries
	ads advertisers.AdvertiserService
}

func (s AdvertiserController) Load(e *echo.Echo) error {
	e.GET("/api/v1/advertisers", s.ListAdvertisers())
	return nil
}

func (s AdvertiserController) ListAdvertisers() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req advertisers.ListAdvertisersInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := advertisers.ListAdvertisersUseCase(c.Request().Context(), req, s.db)
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
		out, err := s.ads.Create(c.Request().Context(), req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}
