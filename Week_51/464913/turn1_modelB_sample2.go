package main

import (
	"fmt"
	"net/url"
)

func main() {

	// Encoded query parameters string
	encodedParams := "name=Abak&age=30&location=New+York&hobbies=coding&hobbies=reading"

	// Create a url.Values type params to store decoded query parameters
	// Decode the encoded query parameters
	params, err := url.ParseQuery(encodedParams)
	if err != nil {
		fmt.Println("Error decoding query parameters:", err)
		return
	}

	// Print the decoded query parameters
	fmt.Println(params)
	// Output: map[hobbies:[coding reading] location:[New York] name:[Abak] age:[30]]
}
