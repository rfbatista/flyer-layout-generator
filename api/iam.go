package api

import (
	"github.com/labstack/echo/v4"
)

func NewIAMAPI() IAMController {
	return IAMController{}
}

type IAMController struct{}

func (i IAMController) Load(e *echo.Echo) error {
	return nil
}
