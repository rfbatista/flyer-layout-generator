package api

import (
	"algvisual/internal/clients"
	"algvisual/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewClientsController(db *database.Queries) ClientController {
	return ClientController{db: db}
}

type ClientController struct {
	db *database.Queries
}

func (s ClientController) Load(e *echo.Echo) error {
	e.GET("/api/v1/clients", s.ListClients())
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
