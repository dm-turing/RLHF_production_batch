package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Map function to square each number
	squares := make([]int, len(numbers))
	map(func(i int) { squares[i] = numbers[i] * numbers[i] }, numbers)

	// Reduce function to sum the squares
	sumOfSquares := 0
	for _, square := range squares {
		sumOfSquares += square
	}

	fmt.Println("Sum of squares:", sumOfSquares)
}