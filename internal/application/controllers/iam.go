package controllers

import (
	"algvisual/internal/infrastructure/cognito"
	"algvisual/internal/infrastructure/config"
	"algvisual/internal/infrastructure/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewIAMController(
	cog *cognito.Cognito,
	c config.AppConfig,
) IAMController {
	return IAMController{cog: cog, c: c}
}

type IAMController struct {
	cog *cognito.Cognito
	c   config.AppConfig
}

func (i IAMController) Load(e *echo.Echo) error {
	e.GET(
		"/whoami",
		i.WhoAmI(),
		middlewares.NewAuthMiddleware(i.cog, i.c),
	)
	return nil
}

func (i IAMController) WhoAmI() echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*middlewares.ApplicationContext)
		return c.JSON(http.StatusOK, cc.UserSession())
	}
}
