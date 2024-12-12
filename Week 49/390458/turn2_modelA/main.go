package main

import (
	"fmt"
	"log"

	"github.com/example/models" // Update to your path
	"github.com/example/parser" // Update to your path
)

func main() {
	err := parser.StreamParseFile("large_data.xml", func(person models.Person) {
		fmt.Printf("Name: %s, Age: %d, Address: %s\n", person.Name, person.Age, person.Address)
	})

	if err != nil {
		log.Fatalf("Error parsing XML: %v", err)
	}
}
