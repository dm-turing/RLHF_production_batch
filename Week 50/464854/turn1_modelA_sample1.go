package main

import (
	"fmt"
)

func main() {
	var largeSlice []int
	for i := 0; i < 1000000; i++ {
		largeSlice = append(largeSlice, i)
	}
	// The largeSlice is still accessible, so memory is not freed
	fmt.Println("Initial slice size:", cap(largeSlice))
}
