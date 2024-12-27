package main

import (
	"fmt"
	"time"
)

func main() {
	var sum int
	arr := make([]int, 100000000)
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	start := time.Now()
	for _, value := range arr {
		sum += value
	}
	end := time.Now()
	fmt.Println("Sum:", sum)
	fmt.Println("Range statement time:", end.Sub(start))
}
