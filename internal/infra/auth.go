package infra

import (
	"strings"

	"github.com/labstack/echo/v4"
)

func NewAuthMiddleware(cog *Cognito) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			key := c.Request().Header.Get("Authorization")
			err := cog.VerifyToken(
				c.Request().Context(),
				[]byte(strings.Split(key, "Bearer ")[1]),
			)
			if err != nil {
				return err
			}
			return next(c)
		}
	}
}
