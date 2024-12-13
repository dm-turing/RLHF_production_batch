package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"`
}

func main() {
	jsonData := `{"name":"Bob","age":25}`

	var p Person
	err := json.Unmarshal([]byte(jsonData), &p)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %s", err)
	}

	fmt.Printf("Person: %+v\n", p) // Output: Person: {Name:Bob Age:25 Email:}
}
