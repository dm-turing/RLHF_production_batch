package main

import (
	"fmt"
	"net/url"
)

func main() {
	// Example 1: Decoding a single parameter
	encodedParam1 := "Hello%2C%20World%21"
	decodedParam1, err := url.QueryUnescape(encodedParam1)
	if err != nil {
		fmt.Println("Error decoding param1:", err)
	} else {
		fmt.Printf("Decoded param1: %s\n", decodedParam1) // Output: Decoded param1: Hello, World!
	}

	// Example 2: Decoding multiple values for a single key
	encodedValues := "key=value1&key=value2"
	parsedValues, err := url.ParseQuery(encodedValues)
	if err != nil {
		fmt.Println("Error parsing query:", err)
	} else {
		for _, value := range parsedValues["key"] {
			fmt.Println("Decoded value:", value) // Output: Decoded value: value1
			fmt.Println("Decoded value:", value) // Output: Decoded value: value2
		}
	}

	// Example 3: Decoding a structured query
	encodedStruct := "name=John+Doe&age=30&email=john%40example.com"
	parsedStruct, err := url.ParseQuery(encodedStruct)
	if err != nil {
		fmt.Println("Error parsing query:", err)
	} else {
		fmt.Println("Decoded name:", parsedStruct["name"][0])   // Output: Decoded name: John Doe
		fmt.Println("Decoded age:", parsedStruct["age"][0])     // Output: Decoded age: 30
		fmt.Println("Decoded email:", parsedStruct["email"][0]) // Output: Decoded email: john@example.com
	}
}
