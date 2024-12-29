package main

import (
	"fmt"
	"regexp"
	"strings"
)

func normalizeQueryParamsWithRegex(query string) string {
	// Define a regular expression pattern to handle & and other special characters
	pattern := regexp.MustCompile(`(?%26)&(?%26)`)
	query = pattern.ReplaceAllString(query, "&amp;")

	// Handle other characters as needed
	query = strings.ReplaceAll(query, "%3B", ";")
	query = strings.ReplaceAll(query, "%3A", ":")
	query = strings.ReplaceAll(query, "%2C", ",")
	query = strings.ReplaceAll(query, "%20", " ")

	return query
}

func main() {
	query := "param1=value1&param2=value%3B2&param3=value3%20with%20spaces"
	normalizedQuery := normalizeQueryParamsWithRegex(query)
	fmt.Println("Original Query: ", query)
	fmt.Println("Normalized Query: ", normalizedQuery)
}
