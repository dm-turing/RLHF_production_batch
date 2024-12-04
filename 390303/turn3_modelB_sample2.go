package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	var counter int
	var mu sync.Mutex

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	errChan := make(chan error, 2)

	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				select {
				case <-ctx.Done():
					return
				default:
					mu.Lock()
					counter++
					if counter < 0 {
						errChan <- fmt.Errorf("counter went negative")
					}
					mu.Unlock()
				}
			}
			errChan <- nil
		}(i)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}

	fmt.Printf("Final Counter: %d\n", counter)
}
