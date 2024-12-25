package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Model
type Product struct {
	ID    int
	Name  string
	Price float64
}

var products []Product

// View
func renderProduct(w http.ResponseWriter, p Product) {
	fmt.Fprintf(w, "Product ID: %d, Name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
}

// Controller
func getProductHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	r.ParseForm()
	idStr := r.Form.Get("id")

	// Validate and find the product
	if idStr == "" {
		http.Error(w, "Product ID not provided", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var product Product
	for _, p := range products {
		if p.ID == id {
			product = p
			break
		}
	}

	// If product not found, send a 404 response
	if product.ID == 0 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// Render the product
	renderProduct(w, product)
}

func main() {
	// Initialize products
	products = []Product{
		{ID: 1, Name: "Laptop", Price: 1000.0},
		{ID: 2, Name: "Smartphone", Price: 500.0},
		{ID: 3, Name: "Tablet", Price: 300.0},
	}

	// Set up HTTP server
	http.HandleFunc("/product", getProductHandler)
	fmt.Println("Server running on http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
