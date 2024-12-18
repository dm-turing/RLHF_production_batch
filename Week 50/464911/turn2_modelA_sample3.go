package main

import (
	"fmt"
)

// Function that takes two arguments
func add(x, y int) int {
	return x + y
}

// Partially applied function with x fixed to 10
func addTen(y int) int {
	return add(10, y)
}

func main() {
	result := addTen(20)
	fmt.Println("Result:", result)
}
