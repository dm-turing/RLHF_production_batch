package main

import (
	"fmt"
	"sync"
)

// Map function that applies a given function to each element of the slice.
func mapSlice(nums []int, fn func(int) int) []int {
	out := make([]int, len(nums))
	for i, n := range nums {
		out[i] = fn(n)
	}
	return out
}

// ConcurrentReduce applies a function concurrently to accumulate results.
func concurrentReduce(nums []int, fn func(int, int) int, initial int) int {
	var wg sync.WaitGroup
	resultChan := make(chan int)
	for _, num := range nums {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			resultChan <- fn(initial, n)
		}(num)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	result := initial
	for r := range resultChan {
		result = fn(result, r)
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	// Define a simple doubling function
	double := func(n int) int {
		return n * 2
	}

	// Define a sum function
	sum := func(acc, n int) int {
		return acc + n
	}

	// Apply map operation
	doubledNums := mapSlice(nums, double)
	fmt.Println("Doubled Numbers:", doubledNums)

	// Use concurrent reduce to sum the numbers
	total := concurrentReduce(doubledNums, sum, 0)
	fmt.Println("Sum of Doubled Numbers:", total)
}
