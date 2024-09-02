package middlewares

import (
	"algvisual/internal/shared"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type HTTPErrorResult struct {
	Status      string    `json:"status,omitempty"`
	StatusCode  int       `json:"status_code,omitempty"`
	RequestID   string    `json:"request_id,omitempty"`
	DocumentURL string    `json:"document_url,omitempty"`
	Error       HTTPError `json:"error,omitempty"`
}

type HTTPError struct {
	Code       string    `json:"code,omitempty"`
	Message    string    `json:"message,omitempty"`
	Details    string    `json:"details,omitempty"`
	Timestamp  time.Time `json:"timestamp,omitempty"`
	Path       string    `json:"path,omitempty"`
	Suggestion string    `json:"suggestion,omitempty"`
}

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
	switch v := err.(type) {
	case *shared.AppError:
		result.StatusCode = v.StatusCode
		errorDetails.Message = v.Message
		errorDetails.Timestamp = v.Timestamp
		errorDetails.Details = v.Detail
		errorDetails.Code = string(v.ErrorCode)
		result.Error = errorDetails
	case *echo.HTTPError:
		result.StatusCode = v.Code
		errorDetails.Timestamp = time.Now()
		errorDetails.Message = v.Message.(string)
	case *shared.InternalError:
		result.StatusCode = 500
		errorDetails.Timestamp = time.Now()
		errorDetails.Message = v.Message
	default:
		result.StatusCode = 500
		errorDetails.Message = err.Error()
		errorDetails.Timestamp = time.Now()
	}
	result.Error = errorDetails
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
