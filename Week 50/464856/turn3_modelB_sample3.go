package main

import (
	"log"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

// ... (Previous code snippets)

func main() {
	// ... (Previous code snippets)

	// Create a rotating logger
	logWriter, err := rotatelogs.New(
		auditLogFile+".%Y%m%d",                    // Rotate log file every day
		rotatelogs.WithLinkName(auditLogFile),     // Keep a link to the current log file
		rotatelogs.WithMaxAge(7*24*time.Hour),     // Retain log files for 7 days
		rotatelogs.WithRotationTime(24*time.Hour), // Rotate at midnight
	)
	if err != nil {
		log.Fatalf("Error creating rotating logger: %v", err)
	}
	auditLogger := log.New(logWriter, "AUDIT: ", log.LstdFlags)

	// ... (Rest of the code)
}
