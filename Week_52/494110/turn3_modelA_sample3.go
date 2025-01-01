package main

import (
	"fmt"
)

func process(data <-chan int, done chan<- bool) {
	for i := range data {
		fmt.Println("Processing", i)
	}
	done <- true
}

func main() {
	data := make(chan int)
	done := make(chan bool)
	defer close(done)

	go process(data, done)
	go func() {
		for i := 0; i < 10; i++ {
			data <- i
		}
		close(data)
	}()

	<-done
}
