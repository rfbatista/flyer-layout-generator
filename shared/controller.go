package shared

import "github.com/labstack/echo/v4"

type Controller interface {
	Load(e *echo.Echo) error
}
