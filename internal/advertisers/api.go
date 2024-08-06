package advertisers

import (
	"algvisual/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewAdvertiserController(
	db *database.Queries,
	ads AdvertiserService,
) AdvertiserController {
	return AdvertiserController{db: db, ads: ads}
}

type AdvertiserController struct {
	db  *database.Queries
	ads AdvertiserService
}

func (s AdvertiserController) Load(e *echo.Echo) error {
	e.GET("/api/v1/advertisers", s.ListAdvertisers())
	e.POST("/api/v1/advertisers", s.CreateAdvertiser())
	return nil
}

func (s AdvertiserController) ListAdvertisers() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req ListAdvertisersInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := ListAdvertisersUseCase(c.Request().Context(), req, s.db)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s AdvertiserController) CreateAdvertiser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req CreateAdvertiserInput
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
