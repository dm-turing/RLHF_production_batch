package main

import (
	"fmt"
	"net/url"
)

func main() {
	// Example 1: Encoding a single parameter
	param1 := "Hello, World!"
	encodedParam1 := url.QueryEscape(param1)
	fmt.Printf("Encoded param1: %s\n", encodedParam1) // Output: Encoded param1: Hello%2C%20World%21

	// Example 2: Encoding multiple values for a single key using url.Values
	values := url.Values{}
	values.Add("key", "value1")
	values.Add("key", "value2")
	encodedValues := values.Encode()
	fmt.Printf("Encoded values: %s\n", encodedValues) // Output: Encoded values: key=value1&key=value2

	// Example 3: Encoding a structured query
	structValues := url.Values{}
	structValues.Set("name", "John Doe")
	structValues.Set("age", "30")
	structValues.Set("email", "john@example.com")
	encodedStruct := structValues.Encode()
	fmt.Printf("Encoded structured query: %s\n", encodedStruct) // Output: Encoded structured query: name=John+Doe&age=30&email=john%40example.com
}
