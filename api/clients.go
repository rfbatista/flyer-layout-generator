package api

import (
	"algvisual/database"
	"algvisual/internal/clients"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewClientsController(db *database.Queries, sc clients.ClientService) ClientController {
	return ClientController{db: db, sc: sc}
}

type ClientController struct {
	db *database.Queries
	sc clients.ClientService
}

func (s ClientController) Load(e *echo.Echo) error {
	e.GET("/api/v1/clients", s.ListClients())
	e.POST("/api/v1/clients", s.CreateClient())
	return nil
}

func (s ClientController) ListClients() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req clients.ListClientsInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := clients.ListClientsUseCase(c.Request().Context(), req, s.db)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s ClientController) CreateClient() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req clients.CreateClientInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := s.sc.CreateClient(c.Request().Context(), req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}
