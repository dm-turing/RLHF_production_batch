package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Item struct {
	ID        int    `gorm:"primary_key"`
	Name      string `gorm:"size:255"`
	CreatedAt string
}

type Pagination struct {
	Cursor string `json:"cursor"`
	Limit  int    `json:"limit"`
}

func main() {
	db, err := gorm.Open("sqlite3", "./example.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&Item{})

	// Seed the database with some sample data
	if err := db.Create(&Item{Name: "Item 1"}).Error; err != nil {
		log.Fatal(err)
	}
	// ... and more items

	// Example of cursor-based pagination
	pagination := Pagination{
		Limit: 2, // Set the page size
	}

	items, nextCursor, err := getItems(db, pagination)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Current Page:")
	for _, item := range items {
		fmt.Printf("Item ID: %d, Name: %s\n", item.ID, item.Name)
	}

	if nextCursor != "" {
		fmt.Println("\nNext Cursor:", nextCursor)
	} else {
		fmt.Println("\nNo more pages.")
	}
}

func getItems(db *gorm.DB, pagination Pagination) ([]Item, string, error) {
	var items []Item
	query := db.Order("id ASC")

	if pagination.Cursor != "" {
		// Use the cursor to start fetching from the next record
		query = query.Where("id > ?", pagination.Cursor)
	}

	if err := query.Limit(pagination.Limit).Find(&items).Error; err != nil {
		return nil, "", err
	}

	var nextCursor string
	if len(items) > 0 {
		// Get the ID of the last item in the current page to use as the cursor for the next page
		nextCursor = fmt.Sprintf("%d", items[len(items)-1].ID)
	}

	return items, nextCursor, nil
}
