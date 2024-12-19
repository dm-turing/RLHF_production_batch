package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	// Iterate over the channel without using a buffer
	for x := range ch {
		fmt.Println(x)
	}

	// The following line will never be executed due to the range loop blocking until the channel is closed
	fmt.Println("Reached here!")
}
