package main

import (
	"fmt"
	"time"
)

func main() {
	// Unbounded slice append
	var unboundedSlice []int
	for i := 0; i < 1e6; i++ {
		unboundedSlice = append(unboundedSlice, i)
	}

	// Storing large slice in an unreferenced variable
	largeSlice := make([]int, 1e7)
	fmt.Println(len(largeSlice))
	defer func() {
		// This defer statement won't be executed until the main function returns
		// So, the large slice will remain in memory until then
	}()

	// Sharing mutable slice between goroutines without synchronization
	mutableSlice := make([]int, 10)
	go func() {
		for i := 0; i < len(mutableSlice); i++ {
			mutableSlice[i] = i
		}
	}()
	go func() {
		for i := 0; i < len(mutableSlice); i++ {
			fmt.Println(mutableSlice[i])
		}
	}()
	time.Sleep(time.Second)
}
