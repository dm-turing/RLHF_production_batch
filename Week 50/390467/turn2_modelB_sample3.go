package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

// NotFoundError is used for 404 Not Found errors.
type NotFoundError struct {
	message string
}

func (e *NotFoundError) Error() string {
	return e.message
}

func (e *NotFoundError) Code() int {
	return http.StatusNotFound
}

// BadRequestError is used for 400 Bad Request errors.
type BadRequestError struct {
	message string
}

func (e *BadRequestError) Error() string {
	return e.message
}

func (e *BadRequestError) Code() int {
	return http.StatusBadRequest
}

// InternalServerError is used for 500 Internal Server errors.
type InternalServerError struct {
	err error
}

func (e *InternalServerError) Error() string {
	return fmt.Sprintf("internal server error: %v", e.err)
}

func (e *InternalServerError) Code() int {
	return http.StatusInternalServerError
}

// ErrorResponse struct for standard error responses.
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// HandleError responds with a custom error response.
func HandleError(w http.ResponseWriter, r *http.Request, err error) {
	code := err.(interface{ Code() int }).Code()
	msg := err.Error()
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&ErrorResponse{
		Code:    code,
		Message: msg,
	})
}

// ExampleHandler uses custom error types.
func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		HandleError(w, r, &BadRequestError{message: "ID parameter is required"})
		return
	}

	// Simulate an internal error
	HandleError(w, r, &InternalServerError{err: errors.New("database error")})
	return
}

func main() {
	r := chi.NewRouter()
	r.Get("/example", ExampleHandler)

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("ListenAndServe failed: %v", err)
	}
}
