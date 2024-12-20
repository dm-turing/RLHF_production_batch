package main

import (
	"fmt"
	"runtime"
)

func main() {
	expectedSize := 10
	// Example of bounded slice append
	slice := make([]int, 0, expectedSize) // preallocate with capacity
	for i := 0; i < expectedSize; i++ {
		slice = append(slice, i)
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB\n", m.Alloc/1024/1024)
	fmt.Printf("TotalAlloc = %v MiB\n", m.TotalAlloc/1024/1024)
	fmt.Printf("HeapAlloc = %v MiB\n", m.HeapAlloc/1024/1024)
}
