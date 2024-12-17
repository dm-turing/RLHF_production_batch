package main

import (
	"fmt"
	"log"
	"net/http"
)

// HelloWorldHandler is a simple HTTP handler.
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// Here, you could perform any business logic needed
	w.Write([]byte("Hello, World!"))
}

func ValidateHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate presence of a required custom header
		if r.Header.Get("X-Custom-Header") == "" {
			http.Error(w, "Missing X-Custom-Header", http.StatusUnauthorized)
			return
		}
		// If valid, call next handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Create the default router
	mux := http.NewServeMux()

	// Register a handler
	mux.HandleFunc("/", HelloWorldHandler)

	// Wrap the mux with the logging middleware
	loggedMux := ValidateHeaderMiddleware(mux)

	// Start the server
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
