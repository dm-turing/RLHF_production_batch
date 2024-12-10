package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

// Address 1: Data Structure Differences
// Use structs to represent the XML structure, accommodating hierarchical and flat data.

type Person struct {
	XMLName   xml.Name  `xml:"person"`
	Name      string    `xml:"name"`
	Age       int       `xml:"age"`
	Emails    []Email   `xml:"emails>email"`
	Addresses []Address `xml:"addresses>address"`
}

type Email struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

type Address struct {
	Type       string `xml:"type,attr"`
	Street     string `xml:",chardata"`
	City       string `xml:"city"`
	PostalCode string `xml:"postalCode"`
}

// Address 2: Data Types and Validation
// Ensure struct fields align with XML data types.

// Address 3: Namespaces and Qualified Names
// If necessary, handle namespaces by including the namespace in XMLName.
// type Person struct {
// 	XMLName xml.Name `xml:"ns:person"` // Assuming "ns" is the namespace
// }

// Address 4: Data Loss and Integrity
// Implement comprehensive parsing logic to capture all XML elements and attributes.

func xmlToPerson(xmlData []byte) (*Person, error) {
	var p Person
	err := xml.Unmarshal(xmlData, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// Address 5: Encoding and Character Set
// Ensure correct encoding by using UTF-8 for reading/writing.

func readXMLFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// Address 6: Scalability and Performance
// For large files, consider streaming parsing techniques instead of in-memory unmarshalling.

// Address 7: Tooling and Compatibility
// Leverage Go's standard `encoding/xml` and `encoding/json` packages.

func personToJSON(person *Person) ([]byte, error) {
	jsonData, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

// Address 8: Documentation and Maintainability
// Include comments to explain each part of the code and its purpose.

func main() {
	xmlFile := "example.xml"

	xmlData, err := readXMLFile(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML file:", err)
		return
	}

	person, err := xmlToPerson(xmlData)
	if err != nil {
		fmt.Println("Error parsing XML:", err)
		return
	}

	jsonData, err := personToJSON(person)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}
