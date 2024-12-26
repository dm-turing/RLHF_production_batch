package main

import (
	"fmt"
)

func main() {
	// Define an anonymous function that takes two integers and returns their sum
	add := func(a int, b int) int {
		return a + b
	}
	// Call the anonymous function using its variable name
	result := add(10, 20)
	fmt.Println("Sum of 10 and 20 is", result)
}
