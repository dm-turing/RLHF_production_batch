package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

// Define XML structure
type Person struct {
	XMLName xml.Name `xml:"person"`
	ID      string   `xml:"id,attr"` // Attribute
	Name    string   `xml:"name"`    // Element
	Age     int      `xml:"age"`     // Element
	Address Address  `xml:"address"` // Nested Element
}

type Address struct {
	City  string `xml:"city"`
	State string `xml:"state"`
}

func main() {
	// Sample XML Data
	xmlData := `
		<person id="123">
			<name>John Doe</name>
			<age>30</age>
			<address>
				<city>New York</city>
				<state>NY</state>
			</address>
		</person>
	`

	// Unmarshal XML data into struct
	var person Person
	if err := xml.Unmarshal([]byte(xmlData), &person); err != nil {
		log.Fatalf("Error unmarshalling XML: %v", err)
	}

	// Marshal the struct into JSON
	jsonData, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling to JSON: %v", err)
	}

	// Print JSON output
	fmt.Println(string(jsonData))

	// Optionally, write to a file
	if err := os.WriteFile("output.json", jsonData, 0644); err != nil {
		log.Fatalf("Error writing JSON to file: %v", err)
	}
}
