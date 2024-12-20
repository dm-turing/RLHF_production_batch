package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("Heap Allocations: %d\n", m.HeapAlloc)
		time.Sleep(1 * time.Second)
	}
}
