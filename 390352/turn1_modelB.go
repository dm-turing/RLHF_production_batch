package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	file, err := os.OpenFile("example.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Simulate I/O operations
	numOperations := 100
	latencies := make([]time.Duration, 0, numOperations)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numOperations; i++ {
		// Simulate random write operation with latency
		data := make([]byte, 1024)
		rand.Read(data)

		start := time.Now()
		_, err = file.Write(data)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			continue
		}
		end := time.Now()

		latencies = append(latencies, end.Sub(start))
		fmt.Printf("Operation %d: Latency = %v\n", i+1, end.Sub(start))
	}

	// Output latency histogram
	fmt.Println("\nLatency Histogram:")
	for _, latency := range latencies {
		if latency < 1*time.Millisecond {
			fmt.Print(".")
		} else if latency < 10*time.Millisecond {
			fmt.Print("-")
		} else if latency < 100*time.Millisecond {
			fmt.Print("|")
		} else {
			fmt.Print("*")
		}
	}
	fmt.Println()
}
