package main

import (
	"fmt"
	"runtime"
)

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("Escaped to Heap: ", float64(float64(m.HeapReleased/1024)/1024), "MB")
}
