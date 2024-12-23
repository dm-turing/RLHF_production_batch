package main

import (
	"database/sql"
	"fmt"
	"log"
)

func selectUserById(db *sql.DB, userId string) string {
	stmt, err := db.Prepare("SELECT name FROM users WHERE id = ?")
	if err != nil {
		log.Fatalf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	var name string
	if err := stmt.QueryRow(userId).Scan(&name); err != nil {
		log.Fatalf("failed to scan result: %v", err)
	}

	return name
}

func main() {
	// Assume db is a valid *sql.DB connection
	userId := "1234567890"
	userName := selectUserById(new(sql.DB), userId)
	fmt.Println("User name:", userName)
}
