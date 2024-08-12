package middlewares

import (
	"algvisual/internal/infra/cognito"
	"strings"

	"github.com/labstack/echo/v4"
)

func NewAuthMiddleware(cog *cognito.Cognito) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			key := c.Request().Header.Get("Authorization")
			user, err := cog.VerifyToken(
				c.Request().Context(),
				[]byte(strings.Split(key, "Bearer ")[1]),
			)
			if err != nil {
				return err
			}
			cc := c.(*ApplicationContext)
			cc.SetUserSession(*user)
			return next(cc)
		}
	}
}
