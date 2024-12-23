package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func validateURL(urlStr string) (bool, *ErrorResponse) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return false, &ErrorResponse{Code: http.StatusBadRequest, Message: "Invalid URL syntax"}
	}

	if u.Scheme != "https" {
		return false, &ErrorResponse{Code: http.StatusBadRequest, Message: "Unsupported protocol"}
	}

	return true, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	urlStr := r.URL.Query().Get("url")

	isValid, errResp := validateURL(urlStr)
	if errResp != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(errResp.Code)
		json.NewEncoder(w).Encode(errResp)
		log.Printf("Error: %s (Code: %d)", errResp.Message, errResp.Code)
		return
	}

	if isValid {
		fmt.Fprintf(w, "URL is valid: %s", urlStr)
	} else {
		fmt.Fprintf(w, "URL is invalid.")
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	http.HandleFunc("/validate", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
