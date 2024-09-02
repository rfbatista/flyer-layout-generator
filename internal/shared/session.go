package shared

import (
	"algvisual/internal/domain/entities"

	"github.com/labstack/echo/v4"
)

func GetSession(c echo.Context) *entities.UserSession {
	return c.Get("session").(*entities.UserSession)
}
