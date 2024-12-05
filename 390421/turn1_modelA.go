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

	// Create an affinity mask for the first two cores
	// For example, if numCPUs = 4, the mask 0x03 would pin the program to cores 0 and 1.
	affinityMask := 0x03

	// Set the CPU affinity
	runtime.SetAffinity(affinityMask)
	fmt.Println("CPU affinity set to cores:", runtime.GOMAXPROCS(-1)-1)

	// Simulate work
	for i := 0; i < 10; i++ {
		fmt.Println("Working on core:", runtime.GOMAXPROCS(-1)-1)
		time.Sleep(1 * time.Second)
	}
}
