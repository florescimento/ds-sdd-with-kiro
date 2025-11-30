package errors

import (
	"fmt"
	"net/http"
)

// ErrorCode represents a specific error type
type ErrorCode string

const (
	ErrCodeInvalidPayload     ErrorCode = "INVALID_PAYLOAD"
	ErrCodeUnauthorized       ErrorCode = "UNAUTHORIZED"
	ErrCodeNotFound           ErrorCode = "NOT_FOUND"
	ErrCodeConflict           ErrorCode = "CONFLICT"
	ErrCodeRateLimitExceeded  ErrorCode = "RATE_LIMIT_EXCEEDED"
	ErrCodeInternalError      ErrorCode = "INTERNAL_ERROR"
	ErrCodeServiceUnavailable ErrorCode = "SERVICE_UNAVAILABLE"
	ErrCodeGatewayTimeout     ErrorCode = "GATEWAY_TIMEOUT"
)

// AppError represents an application error
type AppError struct {
	Code    ErrorCode      `json:"code"`
	Message string         `json:"message"`
	Details map[string]any `json:"details,omitempty"`
	TraceID string         `json:"trace_id,omitempty"`
	Err     error          `json:"-"`
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (%v)", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// Unwrap returns the wrapped error
func (e *AppError) Unwrap() error {
	return e.Err
}

// HTTPStatus returns the appropriate HTTP status code for the error
func (e *AppError) HTTPStatus() int {
	switch e.Code {
	case ErrCodeInvalidPayload:
		return http.StatusBadRequest
	case ErrCodeUnauthorized:
		return http.StatusUnauthorized
	case ErrCodeNotFound:
		return http.StatusNotFound
	case ErrCodeConflict:
		return http.StatusConflict
	case ErrCodeRateLimitExceeded:
		return http.StatusTooManyRequests
	case ErrCodeServiceUnavailable:
		return http.StatusServiceUnavailable
	case ErrCodeGatewayTimeout:
		return http.StatusGatewayTimeout
	default:
		return http.StatusInternalServerError
	}
}

// NewAppError creates a new application error
func NewAppError(code ErrorCode, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// WithDetails adds details to an error
func (e *AppError) WithDetails(details map[string]any) *AppError {
	e.Details = details
	return e
}

// WithTraceID adds a trace ID to an error
func (e *AppError) WithTraceID(traceID string) *AppError {
	e.TraceID = traceID
	return e
}
