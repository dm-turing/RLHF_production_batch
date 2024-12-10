package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type IORecord struct {
	Operation string
	Duration  time.Duration
	Size      int64
	Timestamp time.Time
}

var ioLogs []IORecord

func logIO(operation string, start time.Time, size int64) {
	duration := time.Since(start)
	ioLogs = append(ioLogs, IORecord{Operation: operation, Duration: duration, Size: size, Timestamp: time.Now()})
}

func readFile(filePath string) {
	start := time.Now()
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}
	logIO("Read", start, int64(len(data)))
}

func writeFile(filePath string, data []byte) {
	start := time.Now()
	err := ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %s", err)
	}
	logIO("Write", start, int64(len(data)))
}

func plotIOMetrics() {
	p := plot.New()

	// Prepare data for plotting
	values := make(plotter.Values, len(ioLogs))
	for i, log := range ioLogs {
		values[i] = float64(log.Duration.Milliseconds())
	}

	bars, err := plotter.NewBarChart(values, 10)
	if err != nil {
		log.Fatal(err)
	}

	p.Add(bars)
	p.Title.Text = "I/O Operations Duration"
	p.X.Label.Text = "Operation"
	p.Y.Label.Text = "Duration (ms)"

	if err := p.Save(5*vg.Inch, 5*vg.Inch, "io_operations.png"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Example of I/O operations
	writeFile("test.txt", []byte("Hello, World!"))
	readFile("test.txt")

	// Plot the I/O metrics
	plotIOMetrics()
	fmt.Println("I/O metrics plotted.")
}
