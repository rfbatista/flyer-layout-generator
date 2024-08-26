package shared

import (
	"errors"
	"fmt"
	"time"
)

func NewError(code ErrorCode, msg, detail string) *AppError {
	return &AppError{
		Message:    msg,
		ErrorCode:  code,
		StatusCode: 400,
		Detail:     detail,
		Timestamp:  time.Now(),
	}
}

func NewAppError(code int, msg string, detail string) *AppError {
	return &AppError{
		Message:    msg,
		StatusCode: code,
		Detail:     detail,
		Timestamp:  time.Now(),
	}
}

type AppError struct {
	StatusCode int
	ErrorCode  ErrorCode
	Inner      error
	// Provide any additional fields that you need.
	Message           string
	AdditionalContext string
	Detail            string
	Timestamp         time.Time
}

// Error is mark the struct as an error.
func (e *AppError) Error() string {
	return fmt.Sprintf(
		"error caused due to %v; message: %v; additional context: %v",
		e.Inner,
		e.Message,
		e.AdditionalContext,
	)
}

// Unwrap is used to make it work with errors.Is, errors.As.
func (e *AppError) Unwrap() error {
	// Return the inner error.
	return e.Inner
}

// WrapWithAppError to easily create a new error which wraps the given error.
func WrapWithAppError(err error, message string, additionalContext string) error {
	details := errors.Join(err, errors.New(message))
	return &AppError{
		Inner:             err,
		Message:           message,
		AdditionalContext: additionalContext,
		Detail:            details.Error(),
		StatusCode:        400,
		Timestamp:         time.Now(),
	}
}
