package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// HelloWorldHandler is a simple HTTP handler.
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// Here, you could perform any business logic needed
	w.Write([]byte("Hello, World!"))
}

// LogRequestMiddleware logs the details of each HTTP request.
func LogRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		r.Header.Set("X-Custom-Header", "")

		// Pass the request to the next handler.
		next.ServeHTTP(w, r)

		// Log the request details.
		duration := time.Since(startTime)
		log.Printf("%s %s %s - Duration: %v", r.Method, r.URL.Path, r.Proto, duration)
	})
}

func main() {
	// Create the default router
	mux := http.NewServeMux()

	// Register a handler
	mux.HandleFunc("/", HelloWorldHandler)

	// Wrap the mux with the logging middleware
	loggedMux := LogRequestMiddleware(mux)

	// Start the server
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
