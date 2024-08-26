package middlewares

import (
	"algvisual/internal/entities"
	"algvisual/internal/infra/cognito"
	"algvisual/internal/infra/config"
	"strings"

	"github.com/labstack/echo/v4"
)

func NewAuthMiddleware(
	cog *cognito.Cognito,
	c config.AppConfig,
) func(echo.HandlerFunc) echo.HandlerFunc {
	if c.APPENV == "local" {
		return func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				cc := c.(*ApplicationContext)
				session := entities.UserSession{
					Username:  "local-user",
					CompanyID: 1,
					UserID:    1,
				}
				cc.SetUserSession(session)
				c.Set("session", session)
				return next(cc)
			}
		}
	}
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
			c.Set("session", user)
			return next(cc)
		}
	}
}
