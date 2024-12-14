package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func producer(ch chan int) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		ch <- i
		time.Sleep(1 * time.Microsecond)
	}
}

func consumer(ch chan int) {
	defer wg.Done()
	for range ch {
		// Simulate work
		time.Sleep(1 * time.Microsecond)
	}
}

func main() {
	ch := make(chan int)
	wg.Add(2)

	go producer(ch)
	go consumer(ch)

	wg.Wait()
	fmt.Println("Unbuffered channel processing complete.")
	close(ch)
}
