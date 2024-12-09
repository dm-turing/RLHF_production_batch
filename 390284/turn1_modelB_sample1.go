package main

import (
	"fmt"
	"time"
)

func sayHello(name string, delay time.Duration) {
	time.Sleep(delay)
	fmt.Println("Hello, ", name)
}

func main() {
	// Start two goroutines that say hello after a delay
	go sayHello("Alice", 1*time.Second)
	go sayHello("Bob", 2*time.Second)

	// Wait for the main goroutine to finish executing all statements
	fmt.Println("Main goroutine is done.")
}
