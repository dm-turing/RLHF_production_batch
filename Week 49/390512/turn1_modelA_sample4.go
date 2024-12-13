package main

import (
	"database/sql"
	"fmt"
)

type DBConfig struct {
	ConnectionString string
}

func connectToDatabase(config *DBConfig) (*sql.DB, error) {
	// Assume a function that tries to connect to a database
	db, err := sql.Open("driver-name", config.ConnectionString)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err)) // Can't recover from this
	}
	return db, nil
}

func main() {
	config := &DBConfig{ConnectionString: ""}
	connectToDatabase(config) // Will panic if it fails to connect
}
