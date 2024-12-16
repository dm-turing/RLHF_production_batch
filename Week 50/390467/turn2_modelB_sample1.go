package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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
	panic(fmt.Errorf("simulated panic"))
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer) // Catch panics and log them
	r.Get("/example", ExampleHandler)

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("ListenAndServe failed: %v", err)
	}
}
