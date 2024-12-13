package main

import (
	"fmt"

	"github.com/tidwall/gjson"
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
	// Example: fetch the name of the first product in the first order
	productName := gjson.Get(jsonData, "customer.orders.0.products.0.name")
	fmt.Println("First Product Name:", productName.String()) // Outputs: Laptop

	// Get total of first order
	total := gjson.Get(jsonData, "customer.orders.0.total")
	fmt.Println("First Order Total:", total.Float()) // Outputs: 150.75
}
