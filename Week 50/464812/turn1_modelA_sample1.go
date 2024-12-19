package main

import "fmt"

func main() {
	fruits := []string{"Apple", "Banana", "Orange"}
	for i, fruit := range fruits {
		fmt.Printf("Index: %d, Fruit: %s\n", i, fruit)
	}
}
