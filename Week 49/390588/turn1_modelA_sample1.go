package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"` // omitempty allows skipping empty fields
}

func main() {
	p := Person{
		Name:  "Alice",
		Age:   30,
		Email: "", // This will be omitted in the output
	}

	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %s", err)
	}

	fmt.Println(string(jsonData)) // Output: {"name":"Alice","age":30}
}
