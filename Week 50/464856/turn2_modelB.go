package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

const (
	auditLogFile = "audit.log"
	bufferSize   = 100
)

func main() {
	// Initialize a logger for audit logs
	logFile, err := os.OpenFile(auditLogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening audit log file: %v", err)
	}
	defer logFile.Close()
	auditLogger := log.New(logFile, "AUDIT: ", log.LstdFlags)

	// Create a buffered channel to collect file system events
	eventChan := make(chan fsnotify.Event, bufferSize)

	// Initialize the file system watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("Error creating file system watcher: %v", err)
	}
	defer watcher.Close()

	// Start a goroutine to process audit events from the buffer
	go handleAuditEvents(auditLogger, eventChan)

	// Add directories to be watched recursively
	err = watcher.Add(".")
	if err != nil {
		log.Fatalf("Error watching directory: %v", err)
	}

	fmt.Println("File system auditing is started...")
	fmt.Println("Ctrl+C to stop.")

	// Read from the watcher and pass events to the buffer
	for {
		select {
		case event := <-watcher.Events:
			eventChan <- event
		case err := <-watcher.Errors:
			auditLogger.Printf("Watcher error: %v", err)
		}
	}
}

// handleAuditEvents reads events from the channel and processes them for auditing
func handleAuditEvents(logger *log.Logger, eventChan chan fsnotify.Event) {
	for event := range eventChan {
		// Add an event type filter if needed
		if event.Op == fsnotify.Create || event.Op == fsnotify.Write || event.Op == fsnotify.Remove || event.Op == fsnotify.Rename {
			logger.Printf("Event: %v - Path: %s", event.Op, event.Name)
		}
	}
}
