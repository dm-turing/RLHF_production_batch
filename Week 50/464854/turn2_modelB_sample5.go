package main

import (
	"fmt"
	"runtime"
)

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("GC Pauses: ", m.PauseTotalNs) // Total nanoseconds paused for GC
}
