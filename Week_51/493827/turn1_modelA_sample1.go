package main

import "fmt"

func modifySliceByValue(s []int) {
	s = append(s, 100) // This creates a new slice with the additional element
}

func main() {
	originalSlice := []int{1, 2, 3}
	fmt.Println("Original slice:", originalSlice)

	modifySliceByValue(originalSlice)

	fmt.Println("Original slice after modification:", originalSlice)
}
