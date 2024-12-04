package main

import (
	"fmt"
	"sync"
)

func concurrentMap(nums []int, fn func(int) int) ([]int, error) {
	var wg sync.WaitGroup
	resultChan := make(chan int)
	errChan := make(chan error)

	for _, num := range nums {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			// if err := fn(n); err != nil {
			// errChan <- err
			// } else {
			resultChan <- fn(n)
			// }
		}(num)
	}

	go func() {
		wg.Wait()
		close(resultChan)
		close(errChan)
	}()

	var results []int
	var err error
	for r := range resultChan {
		results = append(results, r)
	}

	for e := range errChan {
		if e != nil {
			err = e
		}
	}

	return results, err
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	double := func(n int) (r int) {
		r = n * 2
		return
	}

	results, err := concurrentMap(nums, double)
	if err != nil {
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Println("Concurrent Mapping Results:", results)
	}
}
