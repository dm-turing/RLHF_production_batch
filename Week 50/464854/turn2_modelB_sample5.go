package main

import (
	"fmt"
	"runtime"
)

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	data := make(map[int][]int)
	for i := 0; i < 1000; i++ {
		largeSlice := make([]int, i)
		data[i] = largeSlice
		if i == 250 {
			fmt.Println("GC Pauses: ", m.PauseTotalNs) // Total nanoseconds paused for GC
		}
	}
	fmt.Println("GC Pauses: ", m.PauseTotalNs) // Total nanoseconds paused for GC
}
