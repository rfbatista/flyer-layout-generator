package shared

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func RenderComponent(comp templ.Component, c echo.Context) error {
	w := c.Response().Writer
	err := comp.Render(c.Request().Context(), w)
	if err != nil {
		return err
	}
	return nil
}

func InfoNotificationMessage(message string) string {
	return fmt.Sprintf("{\"request-notification\": {\"level\":\"info\",\"message\":\"%s\"}}", message)
}
