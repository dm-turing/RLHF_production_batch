package main

import (
	"fmt"
	"net/http"
	"net/url"

	"go.uber.org/zap"
)

var logger, _ = zap.NewProduction()

func processRequest(w http.ResponseWriter, r *http.Request) {
	urlStr := r.FormValue("url")
	if urlStr == "" {
		http.Error(w, "URL is required.", http.StatusBadRequest)
		logger.Error("Missing URL", zap.String("path", r.URL.Path), zap.String("remote", r.RemoteAddr))
		return
	}

	_, err := url.Parse(urlStr)
	if err != nil {
		http.Error(w, "Invalid URL: "+err.Error(), http.StatusBadRequest)
		logger.Error("Invalid URL", zap.String("path", r.URL.Path), zap.String("remote", r.RemoteAddr), zap.String("url", urlStr), zap.Error(err))
		return
	}

	fmt.Fprintf(w, "Valid URL: %s\n", urlStr)
}

func main() {
	defer logger.Sync()

	http.HandleFunc("/validate", processRequest)
	fmt.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
		logger.Fatal("Server error", zap.Error(err))
	}
}
