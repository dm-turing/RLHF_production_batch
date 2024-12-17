package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
)

type Person struct {
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}

func main() {
	data := `
    <people>
        <person>
            <name>John Doe</name>
            <age>30</age>
        </person>
        <person>
            <name>Jane Doe</name>
            <age>25</age>
        </person>
    </people>`

	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	for {
		// Read tokens from the XML
		t, err := decoder.Token()
		if err != nil {
			break
		}

		// If we have a start element, we decode it into a Person
		if se, ok := t.(xml.StartElement); ok && se.Name.Local == "person" {
			var p Person
			if err := decoder.DecodeElement(&p, &se); err != nil {
				fmt.Fprintf(os.Stderr, "Error decoding element: %v\n", err)
				continue
			}
			fmt.Printf("Parsed Person: %+v\n", p)
		}
	}
}
