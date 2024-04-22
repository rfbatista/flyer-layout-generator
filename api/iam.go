package api

import (
	"github.com/labstack/echo/v4"
	"github.com/rfbatista/apitools"
)

func NewLoginAPI() apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.POST)
	h.SetHandle(func(c echo.Context) error {
		return nil
	})
	return h
}
