package main

import (
	"fmt"
	"net/url"
)

func validateURL(rawurl string) error {
	// Create a URL object from the raw URL string
	u, err := url.Parse(rawurl)
	if err != nil {
		// Return an error if parsing fails
		return fmt.Errorf("Invalid URL format: %w", err)
	}

	// Check if the URL scheme is supported (e.g., http, https)
	if u.Scheme == "" {
		return fmt.Errorf("Invalid URL scheme: %s", rawurl)
	}

	// Validate the hostname component
	if u.Hostname() == "" {
		return fmt.Errorf("Invalid URL hostname: %s", rawurl)
	}
	return nil
}
func main() {
	// Valid URLs
	validURLs := []string{
		"https://www.example.com",
		"http://www.example.org/path/to/resource",
	}

	// Invalid URLs
	invalidURLs := []string{
		"invalidurl", // Missing scheme
		"http://www", // Missing hostname
	}

	// Validate URLs
	for _, url := range validURLs {
		if err := validateURL(url); err != nil {
			fmt.Printf("Invalid URL: %s\n", url)
		} else {
			fmt.Printf("Valid URL: %s\n", url)
		}
	}

	for _, url := range invalidURLs {
		if err := validateURL(url); err != nil {
			fmt.Printf("Invalid URL: %s - %s\n", url, err.Error())
		} else {
			fmt.Printf("Valid URL: %s\n", url)
		}
	}
}
