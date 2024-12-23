package main

import (
	"fmt"
)

type Counter struct {
	value int
}

func (c *Counter) Increment() {
	c.value++
}

func (c *Counter) Decrement() {
	c.value--
}

func processCounter(c *Counter) error {
	c.Increment()
	defer c.Decrement()

	// Simulate some processing
	if c.value > 10 {
		return fmt.Errorf("counter value exceeded limit")
	}
	return nil
}

func main() {
	counter := &Counter{value: 0}
	err := processCounter(counter)
	if err != nil {
		fmt.Printf("Error processing counter: %v\n", err)
	}
	fmt.Println("Final counter value:", counter.value) // Output: Final counter value: 0
}
