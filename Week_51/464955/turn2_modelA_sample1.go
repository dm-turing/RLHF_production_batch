package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func processRequest(w http.ResponseWriter, r *http.Request) {
	urlStr := r.FormValue("url")
	if urlStr == "" {
		http.Error(w, "URL is required.", http.StatusBadRequest)
		return
	}

	_, err := url.Parse(urlStr)
	if err != nil {
		http.Error(w, "Invalid URL: "+err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Valid URL: %s\n", urlStr)
}

func main() {
	http.HandleFunc("/validate", processRequest)
	fmt.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
