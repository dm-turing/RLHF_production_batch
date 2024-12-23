package main

import (
	"fmt"
	"net/url"
	"regexp"
)

func validateURLWithRegex(rawurl string) error {
	regex := regexp.MustCompile(`^(http|https):\/\/[a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,}(\/.*)?$`)
	_, err := url.Parse(rawurl)
	if err != nil {
		return fmt.Errorf("Invalid URL format: %w", err)
	}

	// Check if the URL matches the regular expression pattern
	if !regex.MatchString(rawurl) {
		return fmt.Errorf("Invalid URL format: %s", rawurl)
	}

	// Validate other URL components as before...

	return nil
}
func main() {
	// Add URLs that might fail without the additional regex validation
	invalidURLsWithPath := []string{
		"http://example.com//path//to/resource", // Extra slash after hostname
	}

	// Validate URLs
	for _, url := range invalidURLsWithPath {
		if err := validateURLWithRegex(url); err != nil {
			fmt.Printf("Invalid URL: %s\n", url)
		} else {
			fmt.Printf("Valid URL: %s\n", url)
		}
	}
}
