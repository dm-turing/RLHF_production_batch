package main

import (
	"fmt"
)

type Counter struct {
	count int
}

func (c *Counter) Increment(value int) {
	c.count += value
}

func processData(data []int, callback func(*Counter) int) int {
	var counter Counter
	total := 0

	for _, value := range data {
		total += callback(&counter)
		fmt.Println(value)
	}

	return total
}

func main() {
	data := []int{1, 2, 3, 4, 5}

	total := processData(data, func(counter *Counter) int {
		counter.Increment(1)
		return counter.count
	})

	fmt.Println("Total count:", total)
}
