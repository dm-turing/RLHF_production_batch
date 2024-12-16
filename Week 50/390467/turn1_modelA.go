package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

// CustomError defines a structure to represent an API error.
type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error returns the error message.
func (e *CustomError) Error() string {
	return e.Message
}

// RespondWithError sends a JSON response with the error details.
func RespondWithError(w http.ResponseWriter, err error) {
	var customErr *CustomError
	if ok := errors.As(err, &customErr); ok {
		w.WriteHeader(customErr.Code)
		json.NewEncoder(w).Encode(customErr)
	} else {
		// For unrecognized errors, return a generic 500 error
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&CustomError{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		})
	}
}

// Example function that might return an error.
func getUser(w http.ResponseWriter, r *http.Request) {
	// Simulate an error
	err := someDatabaseCall() // This function simulates a database call which may fail.
	if err != nil {
		RespondWithError(w, &CustomError{
			Code:    http.StatusNotFound,
			Message: "User not found",
		})
		return
	}

	// Normal response if no error occurred...
}

// Middleware for centralized error handling
func errorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				RespondWithError(w, &CustomError{
					Code:    http.StatusInternalServerError,
					Message: "An unexpected error occurred",
				})
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func someDatabaseCall() error {
	return errors.New("database connection failed")
}

// Example of a main function to tie it all together.
func main() {
	http.Handle("/user", errorHandler(http.HandlerFunc(getUser)))
	http.ListenAndServe(":8080", nil)
}
