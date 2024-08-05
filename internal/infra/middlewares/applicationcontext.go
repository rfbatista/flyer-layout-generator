package middlewares

import (
	"algvisual/internal/entities"

	"github.com/labstack/echo/v4"
)

type ApplicationContext struct {
	echo.Context
}

func (a *ApplicationContext) UserSession() entities.UserSession {
	return entities.UserSession{}
}

func NewApplicationContextMiddleware() func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &ApplicationContext{c}
			return next(cc)
		}
	}
}
