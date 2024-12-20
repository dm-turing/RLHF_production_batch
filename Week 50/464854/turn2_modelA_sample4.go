package main

import (
	"fmt"
	"time"
)

func trackGCPauseDistribution() {
	for {
		debug.GCPause()
		fmt.Printf("GC pause distribution: %dms\n", debug.GCPause.Average()/1000000)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	trackGCPauseDistribution()
}
