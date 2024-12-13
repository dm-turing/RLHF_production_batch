package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type Order struct {
	ID       string    `json:"id"`
	Total    float64   `json:"total"`
	Products []Product `json:"products"`
}

type Customer struct {
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Orders []Order `json:"orders"`
}

type Store struct {
	Customer Customer `json:"customer"`
}

func main() {
	// Raw JSON string (can also be read from an API or a file)
	jsonData := `{
		"customer": {
			"name": "John Doe",
			"email": "john@example.com",
			"orders": [
				{
					"id": "123",
					"total": 150.75,
					"products": [
						{
							"id": "p001",
							"name": "Laptop",
							"price": 1000.00,
							"quantity": 1
						},
						{
							"id": "p002",
							"name": "Mouse",
							"price": 20.00,
							"quantity": 2
						}
					]
				},
				{
					"id": "124",
					"total": 25.00,
					"products": [
						{
							"id": "p003",
							"name": "Keyboard",
							"price": 25.00,
							"quantity": 1
						}
					]
				}
			]
		}
	}`

	var store Store
	err := json.Unmarshal([]byte(jsonData), &store)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	fmt.Printf("Customer: %+v\n", store.Customer)
}
