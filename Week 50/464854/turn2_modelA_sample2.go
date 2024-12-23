package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

func trackLiveObjects() {
	for {
		var gcStats debug.GCStats
		debug.ReadGCStats(&gcStats)

		fmt.Printf("Live objects: %d\n", gcStats.LiveObjects)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	trackLiveObjects()
}
