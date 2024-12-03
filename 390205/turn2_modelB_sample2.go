package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var sum int64
	var wg sync.WaitGroup

	const numIncrements = 1000
	const numGoroutines = 10

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			for j := 0; j < numIncrements; j++ {
				atomic.AddInt64(&sum, int64(j))
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("Final sum: %d\n", sum)
}
