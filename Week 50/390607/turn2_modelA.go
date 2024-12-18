package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

const proxyURL = "https://proxy.golang.org" // Default proxy URL for Go modules

func main() {
	log.Println("Starting Go Module Proxy...")

	http.HandleFunc("/v1/mod/", func(w http.ResponseWriter, r *http.Request) {
		transport := &http.Transport{
			DisableKeepAlives: true, // Disable keep-alives to prevent connection problems
		}

		proxy := httputil.NewSingleHostReverseProxy(proxyURL)

		proxy.Transport = transport
		proxy.ServeHTTP(w, r)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
