package main

import "fmt"

func assert(condition bool, message string) {
	if !condition {
		panic(message) // Panic if the assertion fails
	}
}

func main() {
	value := -1
	assert(value >= 0, "value must be non-negative") // Will panic if value < 0

	fmt.Println("Value is:", value) // This line won't execute if panic occurs
}
