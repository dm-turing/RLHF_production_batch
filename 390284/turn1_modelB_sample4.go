package main

import (
	"fmt"
	"time"
)

func worker(id int, ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- id * i
		time.Sleep(time.Duration(id) * 100 * time.Millisecond)
	}
	close(ch)
}

func main() {
	ch := make(chan int)

	go worker(1, ch)
	go worker(2, ch)
	go worker(3, ch)

	for result := range ch {
		fmt.Println("Result:", result)
	}

	fmt.Println("Main goroutine is done.")
}
