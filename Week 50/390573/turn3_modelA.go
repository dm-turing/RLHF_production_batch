package main

import (
	"encoding/xml"
	"log"
	"net/http"
	"strconv"
)

type Item struct {
	ID   int               `xml:"id"`
	Name string            `xml:"name"`
	Tags []string          `xml:"tags"`
	Attr map[string]string `xml:",any,attr"`
}

type Page struct {
	XMLName     xml.Name `xml:"page"`
	Items       []Item   `xml:"item"`
	CurrentPage int      `xml:"current_page"`
	TotalPages  int      `xml:"total_pages"`
	TotalItems  int      `xml:"total_items"`
	TotalTags   int      `xml:"total_tags"`
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	// Example data
	items := []Item{
		{
			ID:   1,
			Name: "Apple",
			Tags: []string{"fruit", "red"},
			Attr: map[string]string{"organic": "true"},
		},
		{
			ID:   2,
			Name: "Banana",
			Tags: []string{"fruit", "yellow"},
		},
		{
			ID:   3,
			Name: "Orange",
			Tags: []string{"fruit", "orange"},
			Attr: map[string]string{"citrus": "true"},
		},
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
	totalTags := 0

	// Slice the items for the requested page
	startIndex := (page - 1) * size
	endIndex := min(totalItems, startIndex+size)
	pagedItems := items[startIndex:endIndex]

	for _, item := range pagedItems {
		totalTags += len(item.Tags)
	}

	// Create the page response
	pageResponse := Page{
		Items:       pagedItems,
		CurrentPage: page,
		TotalPages:  totalPages,
		TotalItems:  totalItems,
		TotalTags:   totalTags,
	}

	// Set content type header and encode the response
	w.Header().Set("Content-Type", "application/xml")
	enc := xml.NewEncoder(w)
	defer enc.Close()

	if err := enc.Encode(pageResponse); err != nil {
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
