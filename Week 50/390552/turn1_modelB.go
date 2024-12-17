package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware is a custom middleware that logs the start and end of each request
func LoggingMiddleware(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[INFO] %s %s\n", r.Method, r.URL.RequestURI())

		// Call the next handler
		// next.ServeHTTP(w, r)

		log.Printf("[INFO] %s %s %s\n", r.Method, r.URL.RequestURI(), time.Since(start))
	}
}

// AuthenticationMiddleware is a custom middleware for authentication
func AuthenticationMiddleware(next func(http.ResponseWriter, *http.Request), token string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || authHeader != "Bearer "+token {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Call the next handler if authenticated
		// next.ServeHTTP(w, r)
	}
}

// ExampleHandler is a simple handler to demonstrate middleware usage
func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	// Define a simple router
	http.HandleFunc("/", LoggingMiddleware(
		AuthenticationMiddleware(
			ExampleHandler,
			"your-secret-token",
		),
	))

	// Start the server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
