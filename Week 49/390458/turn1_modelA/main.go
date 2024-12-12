package main

import (
	"fmt"
	"log"

	"github.com/example/parser" // Ensure this matches your path
)

func main() {
	people, err := parser.ParseFile("data.xml")
	if err != nil {
		log.Fatalf("Error parsing XML: %v", err)
	}

	for _, person := range people.Persons {
		fmt.Printf("Name: %s, Age: %d, Address: %s\n", person.Name, person.Age, person.Address)
	}
}
