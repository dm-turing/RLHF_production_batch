package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// User represents the structure of our user data
type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

// Extract function to read data from a CSV file
func Extract(filePath string) ([]User, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var users []User
	for _, record := range records[1:] { // Skip header
		id, _ := strconv.Atoi(record[0])
		age, _ := strconv.Atoi(record[3])
		user := User{
			ID:    id,
			Name:  record[1],
			Email: record[2],
			Age:   age,
		}
		users = append(users, user)
	}

	return users, nil
}

// Transform function to apply transformations to the user data
func Transform(users []User) []User {
	for i := range users {
		// Example transformation: Normalize name to upper case
		users[i].Name = strings.ToUpper(users[i].Name)
	}
	return users
}

// Load function to insert data into the PostgreSQL database
func Load(users []User, db *sql.DB) error {
	for _, user := range users {
		_, err := db.Exec("INSERT INTO users (id, name, email, age) VALUES ($1, $2, $3, $4)", user.ID, user.Name, user.Email, user.Age)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	// Step 1: Extract
	users, err := Extract("sample.csv")
	if err != nil {
		log.Fatalf("Error extracting data: %v", err)
	}

	// Step 2: Transform
	transformedUsers := Transform(users)

	// Step 3: Load
	connStr := "user=username dbname=mydb sslmode=disable" // Adjust with your DB credentials
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	if err := Load(transformedUsers, db); err != nil {
		log.Fatalf("Error loading data: %v", err)
	}

	fmt.Println("ETL process completed successfully!")
}
