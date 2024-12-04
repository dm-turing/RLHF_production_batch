package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := processCounter(&counter, &mu, ctx); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}

func processCounter(counter *int, mu *sync.Mutex, ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			mu.Lock()
			*counter++
			if *counter < 0 {
				return errors.New("counter went negative")
			}
			mu.Unlock()
			time.Sleep(time.Microsecond)
		}
	}
}
