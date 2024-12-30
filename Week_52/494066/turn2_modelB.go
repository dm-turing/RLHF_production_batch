package main

import "fmt"

func formattedPrice(price float64) string {
	return fmt.Sprintf("$%.2f", price)
}

func main() {
	product := struct {
		Name     string
		Price    float64
		Discount bool
	}{
		Name:     "Go Programming Language",
		Price:    49.99,
		Discount: true,
	}
	msg := fmt.Sprintf("%s (%s): %s",
		product.Name,
		product.Discount,
		formattedPrice(product.Price),
	)
	fmt.Println(msg)
}
