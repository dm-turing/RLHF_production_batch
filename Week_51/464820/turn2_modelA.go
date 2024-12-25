package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Model
type Product struct {
	ID       int
	Name     string
	Price    float64
	Archived bool
}

var products []Product

// Archive a product
func archiveProduct(id int) bool {
	for i, p := range products {
		if p.ID == id {
			products[i].Archived = true
			return true
		}
	}
	return false
}

// View
func renderJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

// Controller
func archiveProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	r.ParseForm()
	idsStr := r.Form.Get("ids")

	// Split the comma-separated IDs
	ids := strings.Split(idsStr, ",")

	var success bool
	var failures []int

	// Process each product ID
	for _, idStr := range ids {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			failures = append(failures, id)
			continue
		}

		if !archiveProduct(id) {
			failures = append(failures, id)
		} else {
			success = true
		}
	}

	// Construct response
	response := struct {
		Success  bool
		Failures []int
	}{
		Success:  success,
		Failures: failures,
	}

	renderJSON(w, response)
}

func main() {
	// Initialize products
	products = []Product{
		{ID: 1, Name: "Laptop", Price: 1000.0},
		{ID: 2, Name: "Smartphone", Price: 500.0},
		{ID: 3, Name: "Tablet", Price: 300.0},
	}

	// Set up HTTP server
	http.HandleFunc("/archive", archiveProductsHandler)
	fmt.Println("Server running on http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
