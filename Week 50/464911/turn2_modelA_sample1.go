package main

import (
	"fmt"
)

// Function type for a transformation function
type Transformer func(int) int

// Function that applies a transformer to each element in a slice
func applyTransformer(numbers []int, transformer Transformer) []int {
	return map[int]int(func(x int) int {
		return transformer(x)
	})(numbers)
}

func double(x int) int {
	return x * 2
}

func square(x int) int {
	return x * x
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	doubled := applyTransformer(numbers, double)
	squared := applyTransformer(numbers, square)

	fmt.Println("Doubled numbers:", doubled)
	fmt.Println("Squared numbers:", squared)
}
