package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()
	for val := range ch {
		fmt.Println(val)
	}
}
