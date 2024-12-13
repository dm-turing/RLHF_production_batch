package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}

func main() {
	// Read JSON data from a file
	file, err := os.Open("person.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal the JSON data into a Person struct
	var person Person
	err = json.Unmarshal(byteValue, &person)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Print the decoded Person struct
	fmt.Printf("Name: %s, Age: %d, Email: %s, IsAdmin: %t\n", person.Name, person.Age, person.Email, person.IsAdmin)
}
