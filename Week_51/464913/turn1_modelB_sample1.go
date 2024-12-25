package main

import (
	"fmt"
	"net/url"
)

func main() {
	// Create a url.Values type to store query parameters
	params := url.Values{}

	// Add query parameters
	params.Add("name", "Abak")
	params.Add("age", "30")
	params.Add("location", "New York")

	// Add a duplicate parameter to demonstrate how url.Values handles it
	params.Add("hobbies", "coding")
	params.Add("hobbies", "reading")

	// Encode the query parameters
	encodedParams := params.Encode()

	// Print the encoded query parameters
	fmt.Println(encodedParams)
	// Output: name=Abak&age=30&location=New+York&hobbies=coding&hobbies=reading
}
