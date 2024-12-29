package main

import (
	"fmt"
	"strings"
)

func normalizeQueryParams(query string) string {
	// Define the characters to be replaced and their replacements
	replacements := map[string]string{
		"&amp;": "&",
		"&":     "&amp;",
		"%3B":   ";",
		"%3A":   ":",
		"%2C":   ",",
		"%20":   " ",
	}

	for old, new := range replacements {
		query = strings.ReplaceAll(query, old, new)
	}
	return query
}

func main() {
	query := "param1=value1&param2=value%3B2&param3=value3%20with%20spaces"
	normalizedQuery := normalizeQueryParams(query)
	fmt.Println("Original Query: ", query)
	fmt.Println("Normalized Query: ", normalizedQuery)
}
