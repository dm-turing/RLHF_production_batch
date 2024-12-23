package main

import (
	"fmt"
	"sync"
)

var mutex = &sync.Mutex{}

func lockedOperation() {
	mutex.Lock()
	defer mutex.Unlock()

	// Safe to perform operations on shared resources here
	fmt.Println("Operating on shared resource.")
}

func main() {
	go lockedOperation()
	go lockedOperation()
	// Without defer, unlocking might not occur consistently
}
