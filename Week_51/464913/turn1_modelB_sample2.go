package main

import (
	"fmt"
	"net/url"
)

func main() {
	// Create a url.Values type to store decoded query parameters
	// params := url.Values{}

	// Encoded query parameters string
	encodedParams := "name=Abak&age=30&location=New+York&hobbies=coding&hobbies=reading"

	// Decode the encoded query parameters
	parsedValues, err := url.ParseQuery(encodedParams)
	if err != nil {
		fmt.Println("Error decoding query parameters:", err)
		return
	}

	// Print the decoded query parameters
	fmt.Println(parsedValues)
	// Output: map[hobbies:[coding reading] location:[New York] name:[Abak] age:[30]]
}
