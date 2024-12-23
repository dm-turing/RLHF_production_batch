package main

import (
	"fmt"
	"net/url"
)

func validateURLProtocol(urlStr string, allowedProtocols []string) (bool, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return false, fmt.Errorf("Invalid URL syntax: %w", err)
	}
	if !contains(allowedProtocols, u.Scheme) {
		return false, fmt.Errorf("Unsupported protocol: %s", u.Scheme)
	}
	return true, nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	urlStr := "http://example.com"
	allowedProtocols := []string{"http", "https"}
	isValid, err := validateURLProtocol(urlStr, allowedProtocols)
	if err != nil {
		fmt.Println(err)
	} else if isValid {
		fmt.Println("URL protocol is valid.")
	} else {
		fmt.Println("URL protocol is invalid.")
	}
}
