package main

import (
	"fmt"
	"regexp"
)

func normalizeQueryParams(query string) string {
	// Define a slice of replacement rules using regular expressions
	replacements := []struct {
		pattern     *regexp.Regexp
		replacement string
	}{
		{regexp.MustCompile(`&amp;`), "&"},
		{regexp.MustCompile(`&`), "&amp;"},
		{regexp.MustCompile(`%3B`), ";"},
		{regexp.MustCompile(`%3A`), ":"},
		{regexp.MustCompile(`%2C`), ","},
		{regexp.MustCompile(`%20`), " "},
		{regexp.MustCompile(`\+`), " "},                        // Add additional rules for '+' sign
		{regexp.MustCompile(`(?i)[^a-zA-Z0-9_.\-=]`), "%25$1"}, // URL-encode special characters
	}

	for _, rule := range replacements {
		query = rule.pattern.ReplaceAllString(query, rule.replacement)
	}

	return query
}

func main() {
	query := "param1=value1&param2=value%3B2&param3=value3+with+spaces&param4=some_special_characters!@#$%^&*()"
	normalizedQuery := normalizeQueryParams(query)
	fmt.Println("Original Query: ", query)
	fmt.Println("Normalized Query: ", normalizedQuery)
}
