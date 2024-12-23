package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var cache sync.Map

	// Function to simulate fetching data from a remote source
	fetchData := func(key string) string {
		time.Sleep(10 * time.Millisecond) // Simulate slow remote call
		return fmt.Sprintf("Data for key %s", key)
	}

	// Start multiple goroutines to access the cache
	for i := 0; i < 100; i++ {
		go func() {
			key := fmt.Sprintf("key-%d", i)
			if value, ok := cache.Load(key); ok {
				fmt.Println(fmt.Sprintf("Cached data for %s: %s", key, value.(string)))
			} else {
				value := fetchData(key)
				cache.Store(key, value)
				fmt.Println(fmt.Sprintf("Fetched and cached data for %s: %s", key, value))
			}
		}()
	}

	// Wait for all goroutines to finish
	time.Sleep(1 * time.Second)
}
