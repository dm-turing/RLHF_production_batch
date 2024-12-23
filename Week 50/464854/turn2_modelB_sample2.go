package main

import (
	"fmt"
	"runtime"
)

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("Live Slice Count: ", float64(float64(m.HeapObjects/1024)/1024)) // Approximate slice count
}
