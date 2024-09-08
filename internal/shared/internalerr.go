package shared

import (
	"fmt"
	"time"
)

func NewInternalError(msg string) *AppError {
	return &AppError{
		Message:    msg,
		StatusCode: 500,
		Timestamp:  time.Now(),
	}
}

func NewInternalErrorWithDetails(msg string, detail string) *AppError {
	return &AppError{
		Message:    msg,
		StatusCode: 500,
		Detail:     detail,
		Timestamp:  time.Now(),
	}
}

type InternalError struct {
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
func (e *InternalError) Error() string {
	return fmt.Sprintf(
		"error caused due to %v; message: %v; additional context: %v",
		e.Inner,
		e.Message,
		e.AdditionalContext,
	)
}

// Unwrap is used to make it work with errors.Is, errors.As.
func (e *InternalError) Unwrap() error {
	// Return the inner error.
	return e.Inner
}
