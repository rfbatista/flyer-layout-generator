package middlewares

import (
	"algvisual/internal/domain/entities"

	"github.com/labstack/echo/v4"
)

type ApplicationContextKey string

const (
	ApplicationContextKeyUser ApplicationContextKey = "user"
)

type ApplicationContext struct {
	echo.Context
}

func (a *ApplicationContext) UserSession() entities.UserSession {
	return a.Get(string(ApplicationContextKeyUser)).(entities.UserSession)
}

func (a *ApplicationContext) SetUserSession(u entities.UserSession) {
	a.Set(
		string(ApplicationContextKeyUser), u,
	)
}

func NewApplicationContextMiddleware() func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &ApplicationContext{c}
			return next(cc)
		}
	}
}
