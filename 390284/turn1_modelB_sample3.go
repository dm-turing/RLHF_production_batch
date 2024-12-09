package main

import (
	"fmt"
	"sync"
)

var count int
var mux sync.Mutex

func increment(wg *sync.WaitGroup) {
	wg.Done()
	for i := 0; i < 100000; i++ {
		mux.Lock()
		count++
		mux.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10) // Create 10 goroutines

	for i := 0; i < 10; i++ {
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("Final count:", count)
}
