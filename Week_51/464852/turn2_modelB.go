package main

// Import the required package
import (
	"fmt"
	"time"

	"github.com/rcrowley/go-metrics"
)

// Initialize a metrics registry
var registry = metrics.NewRegistry()

func main() {
	// Create a timer metric
	timer := metrics.NewTimer()
	registry.Register("timer", timer)

	// Record the time taken for a specific operation
	timer.Time(func() {
		// Perform the operation here
		time.Sleep(5 * time.Second)
	})

	// Report the metrics (e.g., average, rate)
	fmt.Printf("Average time: %v\n", timer.Mean())
	fmt.Printf("Throughput: %v\n", timer.Rate1())
}
