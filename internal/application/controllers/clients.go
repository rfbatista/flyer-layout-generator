package controllers

import (
	"algvisual/internal/application/usecases/clients"
	"algvisual/internal/infrastructure/cognito"
	"algvisual/internal/infrastructure/config"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewClientsController(
	db *database.Queries,
	sc clients.ClientService,
	cfg config.AppConfig,
	cog *cognito.Cognito,
) ClientController {
	return ClientController{
		db:  db,
		sc:  sc,
		cog: cog,
		cfg: cfg,
	}
}

type ClientController struct {
	db  *database.Queries
	sc  clients.ClientService
	cfg config.AppConfig
	cog *cognito.Cognito
}

func (s ClientController) Load(e *echo.Echo) error {
	e.GET("/api/v1/clients",
		s.ListClients(),
		middlewares.NewAuthMiddleware(s.cog, s.cfg),
	)
	e.POST(
		"/api/v1/clients",
		s.CreateClient(),
		middlewares.NewAuthMiddleware(s.cog, s.cfg),
	)
	return nil
}

func (s ClientController) ListClients() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req clients.ListClientsInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := s.sc.ListClients(c, req)
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
		out, err := s.sc.CreateClient(c, req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}
