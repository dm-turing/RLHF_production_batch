package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
}

func main() {
	data := `
    <person>
        <name>John Doe</name>
        <age>30</age>
    </person>`

	var p Person
	err := xml.Unmarshal([]byte(data), &p)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshalling XML: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Parsed: %+v\n", p)
}
