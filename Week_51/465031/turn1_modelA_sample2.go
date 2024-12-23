package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var counterMap map[string]int = make(map[string]int)
	var mutex sync.Mutex

	// Start multiple goroutines to update the counter
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				counterMap["count"]++
				mutex.Unlock()
			}
		}()
	}

	// Wait for all goroutines to finish
	time.Sleep(1 * time.Second)

	// Retrieve and print the final count
	fmt.Println("Total requests:", counterMap["count"])
}
