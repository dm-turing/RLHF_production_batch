package main

import (
	"fmt"
)

func modifySlice(s []int) {
	s[0] = 100
}

func main() {
	originalSlice := []int{1, 2, 3}
	fmt.Println("Original slice:", originalSlice) // Output: Original slice: [1 2 3]
	modifySlice(originalSlice)
	fmt.Println("Original slice after modification:", originalSlice) // Output: Original slice after modification: [100 2 3]
}
