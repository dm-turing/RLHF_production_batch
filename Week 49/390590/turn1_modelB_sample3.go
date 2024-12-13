package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		ch1 <- "Message from ch1"
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		ch2 <- "Message from ch2"
	}()

	select {
	case msg := <-ch1:
		fmt.Println("Received from ch1:", msg)
	case msg := <-ch2:
		fmt.Println("Received from ch2:", msg)
	default:
		fmt.Println("Timed out")
		time.Sleep(100 * time.Millisecond) // Continue loop
	}
}
