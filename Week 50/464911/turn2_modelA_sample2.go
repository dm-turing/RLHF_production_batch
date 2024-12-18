package main

import (
	"fmt"
)

type Pipeline func(int) int

// Define pipeline functions
func increment(x int) int {
	return x + 1
}

func decrement(x int) int {
	return x - 1
}

func multiplyByTwo(x int) int {
	return x * 2
}

func main() {
	value := 3
	pipeline := multiplyByTwo
	pipeline = func(x int) int {
		return pipeline(decrement(increment(x)))
	}

	result := pipeline(value)
	fmt.Println("Result:", result)
}
