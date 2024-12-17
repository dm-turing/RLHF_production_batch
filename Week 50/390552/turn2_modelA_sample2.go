package main

import (
	"context"
	"fmt"
	"net/http"
)

// Key for the request context
type contextKey string

const userIDKey contextKey = "userID"

// RequestModifier Middleware
func RequestModifier(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add a new value to the request context
		userID := "12345" // This value could come from a token or database
		ctx := context.WithValue(r.Context(), userIDKey, userID)
		r = r.WithContext(ctx)

		// Proceed to the next handler
		next.ServeHTTP(w, r)
	})
}

// Example handler to retrieve the userID from context
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(userIDKey).(string)
	w.Write([]byte(fmt.Sprintf("Hello, User ID: %s!", userID)))
}

// CORS Middleware
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                              // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")            // Allowed methods
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Custom-Header") // Allowed headers

		if r.Method == http.MethodOptions {
			// Preflight request
			w.WriteHeader(http.StatusOK)
			return
		}

		// Proceed to the next handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloWorldHandler)

	// Wrap with middleware
	loggedMux := CORSMiddleware(RequestModifier(mux))

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		fmt.Printf("Could not start server: %s\n", err)
	}
}
