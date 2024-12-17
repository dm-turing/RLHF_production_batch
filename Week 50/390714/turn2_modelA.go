package main

import (
	"fmt"
	"time"
)

// Simulate the checking health of the database.
func isDatabaseHealthy(dbType string) (bool, string) {
	if dbType == "primary" && time.Now().Hour()%2 == 0 {
		return false, "Primary database is down"
	}
	return true, "Database is healthy"
}

// Simulate the successful retrieval of data from a database.
func getData(dbType string) (string, error) {
	if dbType == "primary" {
		return "Data from primary", nil
	} else if dbType == "secondary" {
		return "Data from secondary", nil
	}
	return "", fmt.Errorf("Unknown database: %s", dbType)
}

func handleDatabaseFailure() {
	fmt.Println("Primary database is unavailable.")

	// Use a decision tree to determine the fallback strategy.
	switch handleSecondaryDB() {
	case "secondary_successful":
		fmt.Println("Secondary database connection established successfully.")
	case "secondary_failed":
		fmt.Println("Secondary database failed. Shutting down the application.")
	default:
		fmt.Println("Unknown fallback outcome.")
	}
}

func handleSecondaryDB() string {
	secondaryDBHealthy, _ := isDatabaseHealthy("secondary")

	if !secondaryDBHealthy {
		return "secondary_failed"
	}

	_, err := getData("secondary")

	if err != nil {
		return "secondary_failed"
	}

	return "secondary_successful"
}

func main() {
	_, err := isDatabaseHealthy("primary")

	if err != "" {
		handleDatabaseFailure()
	} else {
		fmt.Println("Primary database connection successful.")
	}
}
