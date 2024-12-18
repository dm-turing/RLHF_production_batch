package main

import (
	"encoding/xml"
	"log"
	"net/http"
	"strconv"
)

type Item struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type Page struct {
	XMLName     xml.Name `xml:"page"`
	Items       []Item   `xml:"item"`
	CurrentPage int      `xml:"current_page"`
	TotalPages  int      `xml:"total_pages"`
	TotalItems  int      `xml:"total_items"`
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	// Example data
	items := []Item{
		{ID: 1, Name: "Apple"},
		{ID: 2, Name: "Banana"},
		{ID: 3, Name: "Orange"},
		{ID: 4, Name: "Mango"},
		{ID: 5, Name: "Pineapple"},
		{ID: 6, Name: "Grape"},
		{ID: 7, Name: "Peach"},
		{ID: 8, Name: "Strawberry"},
		{ID: 9, Name: "Blueberry"},
		{ID: 10, Name: "Watermelon"},
	}

	// Parse query parameters for page and size
	pageParam := r.URL.Query().Get("page")
	sizeParam := r.URL.Query().Get("size")

	page, err := strconv.Atoi(pageParam)
	if err != nil {
		http.Error(w, "Invalid page parameter", http.StatusBadRequest)
		return
	}

	size, err := strconv.Atoi(sizeParam)
	if err != nil {
		http.Error(w, "Invalid size parameter", http.StatusBadRequest)
		return
	}

	// Default values
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}

	// Calculate total pages and total items
	totalItems := len(items)
	totalPages := (totalItems + size - 1) / size

	// Slice the items for the requested page
	startIndex := (page - 1) * size
	endIndex := min(totalItems, startIndex+size)
	pagedItems := items[startIndex:endIndex]

	// Create the page response
	pageResponse := Page{
		Items:       pagedItems,
		CurrentPage: page,
		TotalPages:  totalPages,
		TotalItems:  totalItems,
	}

	// Set content type header and encode the response
	w.Header().Set("Content-Type", "application/xml")
	if err := xml.NewEncoder(w).Encode(pageResponse); err != nil {
		log.Fatal(err)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	http.HandleFunc("/items", itemsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
