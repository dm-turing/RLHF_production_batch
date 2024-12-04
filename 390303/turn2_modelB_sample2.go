package main

import (
	"errors"
	"fmt"
	"sync"
)

func main() {
	var counter int
	var mu sync.Mutex
	done := make(chan error, 2)

	for i := 0; i < 2; i++ {
		go func(id int) {
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				if counter < 0 {
					done <- errors.New("counter went negative")
				}
				mu.Unlock()
			}
			done <- nil
		}(i)
	}

	for i := 0; i < 2; i++ {
		if err := <-done; err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}

	fmt.Printf("Final Counter: %d\n", counter)
}
