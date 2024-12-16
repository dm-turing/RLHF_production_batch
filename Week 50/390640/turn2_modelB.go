package main

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func extract(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func transform(records [][]string) [][]string {
	for i := range records {
		if len(records[i]) > 1 { // Assuming the second column needs transformation
			records[i][1] = strings.ToUpper(records[i][1])
		}
	}
	return records
}

func load(db *sql.DB, tableName string, records [][]string) error {
	query := fmt.Sprintf("INSERT INTO %s (column1, column2) VALUES ($1, $2)", tableName)
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, record := range records {
		if len(record) != 2 {
			log.Println("Skipping record with incorrect number of columns:", record)
			continue
		}
		_, err := stmt.Exec(record[0], record[1])
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	filePath := "input.csv"
	tableName := "etl_data"

	dbConnString := "host=localhost dbname=yourdatabase user=youruser password=yourpassword sslmode=disable"
	db, err := sql.Open("postgres", dbConnString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	records, err := extract(filePath)
	if err != nil {
		log.Fatal(err)
	}

	transformedRecords := transform(records)

	err = load(db, tableName, transformedRecords)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ETL process completed successfully.")
}
