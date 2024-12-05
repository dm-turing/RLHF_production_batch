package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Get the number of logical CPUs
	numCPUs := runtime.NumCPU()
	fmt.Println("Number of logical CPUs:", numCPUs)

	// Simulate work
	for i := 0; i < 10; i++ {
		fmt.Println("Working on core:", runtime.GOMAXPROCS(-1)-1)
		time.Sleep(1 * time.Second)
	}
}
