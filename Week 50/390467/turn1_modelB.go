package main

import (
	"errors"
	"net/http"
	"time"
)

// APIError represents a custom API error
type APIError struct {
	Code        int       `json:"code"`
	Message     string    `json:"message"`
	Detail      string    `json:"detail,omitempty"`
	HTTPStatus  int       `json:"httpStatus"`
	OccurredAt  time.Time `json:"occurredAt"`
	InternalErr error     `json:"-"`
}

// Error implements the error interface
func (e *APIError) Error() string {
	return e.Message
}

// WithDetail adds additional detail to the error
func (e *APIError) WithDetail(detail string) *APIError {
	e.Detail = detail
	return e
}

// NewAPIError creates a new APIError
func NewAPIError(code int, message string, internalErr error) *APIError {
	return &APIError{
		Code:        code,
		Message:     message,
		HTTPStatus:  http.StatusInternalServerError,
		OccurredAt:  time.Now(),
		InternalErr: internalErr,
	}
}

// NewBadRequestError creates a new APIError for bad requests
func NewBadRequestError(message string, internalErr error) *APIError {
	return NewAPIError(1000, message, internalErr).WithHTTPStatus(http.StatusBadRequest)
}

// WithHTTPStatus sets the HTTP status code
func (e *APIError) WithHTTPStatus(status int) *APIError {
	e.HTTPStatus = status
	return e
}

func main() {
	// Example usage
	err := NewBadRequestError("Invalid input data", errors.New("failed to parse JSON"))

	// Log the error with details
	logError(err)

	// Send response to client

	sendErrorResponse(http.ResponseWriter{}, err)
}

// logError logs the APIError with full details
func logError(err error) {
	if e, ok := err.(*APIError); ok {
		println("Error:", e.Code, e.Message, e.Detail, e.HTTPStatus, e.OccurredAt, e.InternalErr)
	} else {
		println("Error:", err.Error())
	}
}

// sendErrorResponse sends an API error as an HTTP response
func sendErrorResponse(w http.ResponseWriter, err error) {
	if e, ok := err.(*APIError); ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(e.HTTPStatus)
		_, err = w.Write([]byte(e.MarshalJSON()))
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
