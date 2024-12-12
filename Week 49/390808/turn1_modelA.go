// Using net/http for RESTful API
package main

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	product := Product{ID: "1", Name: "Widget", Price: 100}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func main() {
	http.HandleFunc("/product", getProduct)
	http.ListenAndServe(":8080", nil)
}
