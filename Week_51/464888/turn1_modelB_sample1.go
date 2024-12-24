package main

import (
	"fmt"
	"net/http"
)

func basicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Basic Auth credentials
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// Validate credentials
		if username != "your_username" || password != "your_password" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// Call the next handler if authentication is successful
		next.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/protected-resource", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the protected resource!")
	})
	http.ListenAndServe(":8080", basicAuthMiddleware(http.DefaultServeMux))
}
