package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var counter sync.Map

	// Start multiple goroutines to update the counter
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				v, _ := counter.LoadOrStore("count", 0)
				counter.Store("count", v.(int)+1)
			}
		}()
	}

	// Wait for all goroutines to finish
	time.Sleep(1 * time.Second)

	// Retrieve and print the final count
	count, _ := counter.Load("count")
	fmt.Println("Total requests:", count.(int))
}
