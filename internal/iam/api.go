package iam

import (
	"algvisual/internal/infra/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewIAMController() IAMController {
	return IAMController{}
}

type IAMController struct{}

func (i IAMController) Load(e *echo.Echo) error {
	e.GET("/whoami", i.WhoAmI())
	return nil
}

func (i IAMController) WhoAmI() echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*middlewares.ApplicationContext)
		return c.JSON(http.StatusOK, cc.UserSession())
	}
}
