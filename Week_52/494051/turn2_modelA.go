package main

import (
	"fmt"
	"sync"
	"time"
)

// Define a struct to hold report data
type ReportData struct {
	Source string
	Value  int
}

// Function to generate a report line
func generateReportLine(data ReportData) string {
	return fmt.Sprintf("Source: %s, Value: %d\n", data.Source, data.Value)
}

func main() {
	var wg sync.WaitGroup
	reportLines := make(chan string, 100) // Bounded buffer with a size of 100

	// Generate some sample data concurrently
	wg.Add(2)
	go generateSampleData("Source A", reportLines, 5, &wg)
	go generateSampleData("Source B", reportLines, 3, &wg)

	// Aggregate and write report lines
	go func() {
		for line := range reportLines {
			fmt.Print(line)
		}
		close(reportLines)
	}()

	wg.Wait()
	time.Sleep(2 * time.Second)
}

// Generate sample data and send it to a channel
func generateSampleData(source string, dataChan chan<- string, numSamples int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < numSamples; i++ {
		value := i + 1
		reportLine := generateReportLine(ReportData{Source: source, Value: value})

		// Non-blocking send
		select {
		case dataChan <- reportLine:
		default:
			// Optionally add backpressure handling or drop data
		}
		time.Sleep(500 * time.Millisecond)
	}
}
