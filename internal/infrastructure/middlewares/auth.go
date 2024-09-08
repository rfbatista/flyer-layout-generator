package middlewares

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/cognito"
	"algvisual/internal/infrastructure/config"
	"algvisual/internal/shared"
	"fmt"
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
				c.Set("session", &session)
				return next(cc)
			}
		}
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			key := c.Request().Header.Get("Authorization")
			parsedKey := strings.Split(key, "Bearer ")
			if len(parsedKey) < 2 {
				return shared.NewError(
					INVALID_AUTH,
					"missing bearer token",
					fmt.Sprintf("received: %s", key),
				)
			}
			token := parsedKey[1]
			user, err := cog.VerifyToken(
				c.Request().Context(),
				[]byte(token),
			)
			if err != nil {
				return &shared.AppError{
					StatusCode: 401,
					ErrorCode:  shared.UNAUTHORIZED,
					Message:    "invalid token",
					Detail:     err.Error(),
				}
			}
			cc := c.(*ApplicationContext)
			cc.SetUserSession(*user)
			c.Set("session", user)
			return next(cc)
		}
	}
}
