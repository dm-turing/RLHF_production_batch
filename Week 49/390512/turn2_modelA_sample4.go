package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func connectToDatabase(dataSource string) *sql.DB {
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err)) // Critical: cannot proceed without a database connection
	}
	return db
}

func main() {
	// Assume this connection string is invalid or the database is unreachable
	db := connectToDatabase("user=invalid dbname=test sslmode=disable")

	// Use the db connection...
	defer db.Close()
}
