package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

func trackGCRuns() {
	for {
		var gcStats debug.GCStats
		debug.ReadGCStats(&gcStats)
		fmt.Printf("GC runs: %d\n", gcStats.NumGC)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	trackGCRuns()
}
