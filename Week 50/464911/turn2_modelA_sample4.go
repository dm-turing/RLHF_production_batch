package main

import (
	"fmt"
)

// Curried function that takes two arguments
func curryAdd(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func main() {
	addFive := curryAdd(5)
	result := addFive(10)
	fmt.Println("Result:", result)
}
