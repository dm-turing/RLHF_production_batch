package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("Hello, ", name)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // We have two goroutines

	go sayHello("Alice", &wg)
	go sayHello("Bob", &wg)

	wg.Wait() // Wait for both goroutines to finish
	fmt.Println("Main goroutine is done.")
}
