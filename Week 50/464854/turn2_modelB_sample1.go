package main

import (
	"fmt"
	"runtime"
)

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("Total Memory Allocated: ", float64(float64(m.HeapAlloc/1024)/1024)) // convert to MB()
}
