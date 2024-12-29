package main

import (
	"fmt"
	"strings"
)

func normalizeURLQueryParameter(param string) string {
	// Example replacement rules for normalizing URL query parameters
	param = strings.ReplaceAll(param, " ", "+")   // Replace spaces with '+'
	param = strings.ReplaceAll(param, "%20", "+") // Replace encoded spaces with '+'
	return param
}

func main() {
	originalParam := "Example URL Query Parameter with spaces"
	normalizedParam := normalizeURLQueryParameter(originalParam)

	fmt.Printf("Original parameter: %q\n", originalParam)
	fmt.Printf("Normalized parameter: %q\n", normalizedParam)
}
