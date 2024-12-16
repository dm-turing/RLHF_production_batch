package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// ErrorResponse struct for standard error responses.
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// HandleError responds with a custom error response.
func HandleError(w http.ResponseWriter, r *http.Request, code int, msg string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&ErrorResponse{
		Code:    code,
		Message: msg,
	})
}

// ExampleHandler simulates an error.
func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	err := fmt.Errorf("internal server error")
	log.Err(err).Msg("An error occurred in ExampleHandler")
}

func main() {
	// Configure zerolog
	log.Output(zerolog.ConsoleWriter{}).
		With().
		Timestamp().
		Str("service", "my-service").
		Logger()

	r := chi.NewRouter()
	r.Use(middleware.Recoverer) // Catch panics and log them
	r.Get("/example", ExampleHandler)

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("ListenAndServe failed: %v", err)
	}
}
