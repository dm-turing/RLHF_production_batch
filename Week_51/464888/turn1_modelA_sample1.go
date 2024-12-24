package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/protected", protectedHandler)
	fmt.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Basic authentication required", http.StatusUnauthorized)
		return
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Basic" {
		http.Error(w, "Invalid basic authentication format", http.StatusUnauthorized)
		return
	}

	payload, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		http.Error(w, "Error decoding base64", http.StatusUnauthorized)
		return
	}
	parts = strings.SplitN(string(payload), ":", 2)
	username, password := parts[0], parts[1]
	if username != "user" || password != "pass" {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Access granted to %s\n", username)
}
