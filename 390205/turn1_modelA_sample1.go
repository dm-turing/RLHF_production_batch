package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var value int32 = 100
	var wg sync.WaitGroup
	const numGoroutines = 10

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for {
				oldValue := atomic.LoadInt32(&value)
				newValue := oldValue + 1
				if atomic.CompareAndSwapInt32(&value, oldValue, newValue) {
					fmt.Printf("Goroutine %d updated value to: %d\n", id, newValue)
					break
				}
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("Final value: %d\n", value)
}
