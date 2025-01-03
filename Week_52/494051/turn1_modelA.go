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

// Function to write report lines to a buffer
func writeToReport(reportBuffer *sync.Mutex, lines []string) {
	reportBuffer.Lock()
	defer reportBuffer.Unlock()
	for _, line := range lines {
		fmt.Print(line)
	}
}

func main() {
	var reportBuffer sync.Mutex
	reportLines := make(chan string)

	// Generate some sample data concurrently
	go generateSampleData("Source A", reportLines, 5)
	go generateSampleData("Source B", reportLines, 3)

	// Aggregate and write report lines
	go func() {
		var lines []string
		for line := range reportLines {
			lines = append(lines, line)
		}
		writeToReport(&reportBuffer, lines)
	}()

	// Wait for all data to be processed
	time.Sleep(2 * time.Second)
}

// Generate sample data and send it to a channel
func generateSampleData(source string, dataChan chan<- string, numSamples int) {
	for i := 0; i < numSamples; i++ {
		value := i + 1
		reportLine := generateReportLine(ReportData{Source: source, Value: value})
		dataChan <- reportLine
		time.Sleep(500 * time.Millisecond)
	}
	close(dataChan)
}
