package main

import (
	"fmt"
)

func modifySliceReference(s *[]int) {
	*s = append(*s, 4)
}

func main() {
	originalSlice := []int{1, 2, 3}
	fmt.Println("Original slice:", originalSlice) // Output: Original slice: [1 2 3]
	modifySliceReference(&originalSlice)
	fmt.Println("Original slice after modification:", originalSlice) // Output: Original slice after modification: [1 2 3 4]
}
