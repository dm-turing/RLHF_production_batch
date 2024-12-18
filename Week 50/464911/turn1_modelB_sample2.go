package main

import (
	"fmt"
	"time"
)

func AsyncOperation(done chan<- bool) {
	// Simulate some asynchronous work
	time.Sleep(2 * time.Second)

	// Signal that the operation is complete
	done <- true
}

func main() {
	done := make(chan bool)

	go AsyncOperation(done)

	// Wait for the asynchronous operation to complete
	<-done

	fmt.Println("Operation complete!")
}
