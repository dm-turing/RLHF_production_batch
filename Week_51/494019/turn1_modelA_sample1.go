package main

import (
	"fmt"
	"time"
)

func main() {
	var sum int
	start := time.Now()
	for i := 0; i < 100000000; i++ {
		sum += i
	}
	end := time.Now()
	fmt.Println("Sum:", sum)
	fmt.Println("Traditional for loop time:", end.Sub(start))
}
