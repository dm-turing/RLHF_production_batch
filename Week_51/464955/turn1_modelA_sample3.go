package main

import (
	"fmt"
	"net/url"
)

func validateURLHost(urlStr string) (bool, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return false, fmt.Errorf("Invalid URL syntax: %w", err)
	}
	if u.Host == "" {
		return false, fmt.Errorf("URL host is missing")
	}
	return true, nil
}

func main() {
	urlStr := "http://example.com"
	isValid, err := validateURLHost(urlStr)
	if err != nil {
		fmt.Println(err)
	} else if isValid {
		fmt.Println("URL host is valid.")
	} else {
		fmt.Println("URL host is invalid.")
	}
}
