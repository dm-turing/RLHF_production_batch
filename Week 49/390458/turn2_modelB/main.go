package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

func streamXML(file string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer f.Close()

	d := xml.NewDecoder(f)
	token, err := d.Token()
	if err != nil {
		log.Fatalf("Error getting first token: %v", err)
	}

	// Loop over all the tokens
	for {
		switch se := token.(type) {
		case xml.StartElement:
			fmt.Printf("StartElement: %+v\n", se)
			if se.Name.Local == "person" {
				var person struct {
					XMLName xml.Name `xml:"person"`
					Name    string   `xml:"name"`
					Age     int      `xml:"age"`
					Address string   `xml:"address"`
				}
				for {
					token, err := d.Token()
					if err != nil {
						if err == io.EOF {
							break
						}
						log.Fatalf("Error decoding token: %v", err)
					}
					switch te := token.(type) {
					case xml.CharData:
						// Concatenate multiple CharData chunks into a single string
						person.Name += string(te)
					case xml.EndElement:
						if te.Name.Local == "person" {
							fmt.Printf("Person: Name: %s, Age: %d, Address: %s\n", person.Name, person.Age, person.Address)
							break
						} else {
							fmt.Printf("EndElement: %+v\n", te)
						}
					}
				}
			}
		case xml.EndElement:
			fmt.Printf("EndElement: %+v\n", token.(xml.EndElement))
		case xml.CharData:
			fmt.Printf("CharData: %s\n", string(token.(xml.CharData)))
		default:
			fmt.Printf("Unknown token: %+v\n", token)
		}
		token, err = d.Token()
		if err != nil && err != io.EOF {
			log.Fatalf("Error reading token: %v", err)
		}
		if err == io.EOF {
			break
		}
	}
}

func main() {
	streamXML("data.xml")
}
