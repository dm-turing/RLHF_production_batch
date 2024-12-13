package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func consumer(ch chan int) {
	for num := range ch {
		fmt.Println("Received:", num)
	}
}

func main() {
	ch := make(chan int, 3) // Buffer size of 3

	go producer(ch)
	go consumer(ch)

	time.Sleep(2 * time.Second)
}
