package main

import (
	"fmt"
	"os"
)

func getDatabaseURL() string {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		panic("DATABASE_URL environment variable is required") // Critical error: without this, the app cannot connect to the db
	}
	return url
}

func main() {
	// Ensure DATABASE_URL is set in the environment
	dbURL := getDatabaseURL() // Will panic if DATABASE_URL is not set
	fmt.Println("Database URL is:", dbURL)
}
