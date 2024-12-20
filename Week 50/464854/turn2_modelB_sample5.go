package main

import (
	"debug"
	"fmt"
)

func main() {
	var m debug.MemStats
	debug.ReadMemStats(&m)
	fmt.Println("GC Pauses: ", m.PauseTotalNs) // Total nanoseconds paused for GC
}
