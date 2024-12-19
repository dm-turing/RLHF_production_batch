package main

import (
	"fmt"
)

// Define an Item struct for simplicity
type Item struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
}

// A Cursor is a unique identifier for the starting point of the next page
type Cursor string

func getItems(pageSize int, cursor Cursor) ([]Item, Cursor, error) {
	// Simulate fetching items from a database
	var items []Item
	nextCursor := Cursor("")

	// Example data (replace with actual database query)
	switch cursor {
	case "":
		items = []Item{
			{ID: "1", Name: "Alice", Age: 30, Salary: 70000},
			{ID: "2", Name: "Bob", Age: 25, Salary: 60000},
			{ID: "3", Name: "Charlie", Age: 35, Salary: 80000},
		}
		nextCursor = "3" // Assuming "3" is the last ID in this simulated dataset
	case "3":
		items = []Item{
			{ID: "4", Name: "David", Age: 40, Salary: 90000},
			{ID: "5", Name: "Eve", Age: 32, Salary: 75000},
		}
		nextCursor = "5"
	case "5":
		items = []Item{
			{ID: "6", Name: "Frank", Age: 45, Salary: 100000},
		}
	default:
		return nil, nextCursor, fmt.Errorf("invalid cursor: %s", cursor)
	}

	return items, nextCursor, nil
}

func main() {
	var cursor Cursor = ""
	pageSize := 2

	for {
		items, nextCursor, err := getItems(pageSize, cursor)
		if err != nil {
			fmt.Println("Error fetching items:", err)
			break
		}

		fmt.Printf("Page with cursor %s:\n", cursor)
		for _, item := range items {
			fmt.Println(item)
		}

		if nextCursor == "" {
			break
		}

		cursor = nextCursor
	}
}
