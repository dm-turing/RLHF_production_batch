package main

import (
	"fmt"
	"time"
)

func rateLimiter(ch chan string, rate time.Duration) {
	for {
		ch <- "tick"
		time.Sleep(rate)
	}
}

func main() {
	ch := make(chan string)

	go rateLimiter(ch, 500*time.Millisecond)

	for i := 0; i < 5; i++ {
		msg := <-ch
		fmt.Println("Received:", msg)
	}
}
