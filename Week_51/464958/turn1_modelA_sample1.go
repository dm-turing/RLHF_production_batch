package main

import (
	"fmt"
	"net/url"
)

func main() {
	userInput := "Hello, World!"
	encodedInput := url.QueryEscape(userInput)
	fmt.Println("Encoded input:", encodedInput) // Output: Hello%2C+World%21
}
