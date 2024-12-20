package main

import (
	"fmt"
)

func main() {
	var largestSliceSize int
	for i := 0; i < 1000; i++ {
		slice := make([]int, i*1000)
		if len(slice) > largestSliceSize {
			largestSliceSize = len(slice)
		}
	}
	fmt.Println("Largest Slice Size: ", largestSliceSize)
}
