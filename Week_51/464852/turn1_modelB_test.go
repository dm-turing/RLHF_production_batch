package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

const (
	// JournalPath is the path to the journal directory
	JournalPath = "/path/to/your/journal"
	// JournalFileSize is the size of each journal file in bytes
	JournalFileSize = 1024 * 1024 // 1 MB
	// NumMessages is the number of messages to write to the journal
	NumMessages = 10000
	// InfluxDBURL is the URL of the InfluxDB instance
	InfluxDBURL = "http://localhost:8086"
	// InfluxDBDatabase is the InfluxDB database name
	InfluxDBDatabase = "mydb"
	// InfluxDBBucket is the InfluxDB bucket name
	InfluxDBBucket = "journal_performance"
)

func BenchmarkJournalPerformance(b *testing.B) {
	// Create an InfluxDB client
	client := influxdb2.NewClient(InfluxDBURL, "your-influxdb-token")
	defer client.Close()
	writeAPI := client.WriteAPIBlocking(InfluxDBDatabase, InfluxDBBucket)
	// defer writeAPI.Close()

	// Create a temporary directory for the journal files
	tempDir, err := ioutil.TempDir("", "journal-benchmark-")
	if err != nil {
		b.Fatalf("Error creating temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Generate random messages
	messages := make([][]byte, NumMessages)
	for i := range messages {
		messages[i] = []byte(fmt.Sprintf("message-%d-%d", i, rand.Intn(1000)))
	}

	// Perform the benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a new journal file
		journalFilePath := filepath.Join(tempDir, fmt.Sprintf("journal-%d.log", i))
		f, err := os.Create(journalFilePath)
		if err != nil {
			b.Fatalf("Error creating journal file: %v", err)
		}
		defer f.Close()

		// Write messages to the journal file
		for _, message := range messages {
			_, err := f.Write(message)
			if err != nil {
				b.Fatalf("Error writing message to journal file: %v", err)
			}
		}

		// Sync the journal file to disk
		err = f.Sync()
		if err != nil {
			b.Fatalf("Error syncing journal file: %v", err)
		}

		// Calculate the journal file size
		fileInfo, err := f.Stat()
		if err != nil {
			b.Fatalf("Error getting journal file info: %v", err)
		}
		journalFileSize := fileInfo.Size()

		// Write performance metrics to InfluxDB
		point := influxdb2.NewPoint("journal_performance",
			map[string]string{"journal_file": journalFilePath},
			map[string]interface{}{
				"message_count":     len(messages),
				"journal_file_size": journalFileSize,
				"write_time":        float64(time.Since(time.Now())) / float64(time.Millisecond),
			},
			time.Now(),
		)
		writeAPI.WritePoint(context.Background(), point)
	}
}
