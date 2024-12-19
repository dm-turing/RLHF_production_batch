package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	for val := range ch {
		fmt.Println(val)
	}
}
