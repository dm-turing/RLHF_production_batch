package main

import (
	"fmt"
)

func main() {
	numbers := make(chan int, 3)
	go func() {
		numbers <- 1
		numbers <- 2
		numbers <- 3
		close(numbers)
	}()

	for num := range numbers {
		fmt.Println(num)
	}
}
