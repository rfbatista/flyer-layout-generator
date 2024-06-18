package infra

import (
	"algvisual/internal/shared"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func ErrorHandler() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	}
}

func HTTPErrorHandler(err error, c echo.Context) {
	var result HTTPErrorResult
	var errorDetails HTTPError
	he, ok := err.(*shared.AppError)
	result.Status = "error"
	if ok {
		result.StatusCode = he.StatusCode
		errorDetails.Message = he.Message
		errorDetails.Timestamp = he.Timestamp
		errorDetails.Details = he.Detail
		result.Error = errorDetails
	} else {
		result.StatusCode = 500
		errorDetails.Message = err.Error()
		errorDetails.Timestamp = time.Now()
	}
	result.Error = errorDetails
	shared.ErrorNotification(c, "Falha interna")
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(result.StatusCode)
		} else {
			err = c.JSON(result.StatusCode, result)
		}
		if err != nil {
			c.Echo().Logger.Error(err)
		}
	}
}
