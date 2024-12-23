package main

import (
	"fmt"
	"net/url"
)

func validateURLSyntax(urlStr string) (bool, error) {
	_, err := url.Parse(urlStr)
	if err != nil {
		return false, fmt.Errorf("Invalid URL syntax: %w", err)
	}
	return true, nil
}

func main() {
	urlStr := "http://example.com"
	isValid, err := validateURLSyntax(urlStr)
	if err != nil {
		fmt.Println(err)
	} else if isValid {
		fmt.Println("URL syntax is valid.")
	} else {
		fmt.Println("URL syntax is invalid.")
	}
}
