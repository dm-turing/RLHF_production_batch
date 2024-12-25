package main

import "fmt"

func modifySliceByReference(s *[]int) {
	*s = append(*s, 100) // This modifies the original slice
}

func main() {
	originalSlice := []int{1, 2, 3}
	fmt.Println("Original slice:", originalSlice)

	modifySliceByReference(&originalSlice)

	fmt.Println("Original slice after modification:", originalSlice)
}
