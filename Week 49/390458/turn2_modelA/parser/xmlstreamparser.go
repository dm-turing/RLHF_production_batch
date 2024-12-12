package parser

import (
	"encoding/xml"
	"io"
	"os"

	"github.com/example/models" // Make sure to update the import path
)

func StreamParseXML(data io.Reader, callback func(models.Person)) error {
	decoder := xml.NewDecoder(data)

	for {
		// Read tokens from the XML
		t, err := decoder.Token()
		if err == io.EOF {
			break // End of the XML
		} else if err != nil {
			return err
		}

		// Check if the token is a start element
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "person" {
				var person models.Person
				if err := decoder.DecodeElement(&person, &se); err != nil {
					return err
				}
				callback(person) // Process the person
			}
		}
	}
	return nil
}

func StreamParseFile(filename string, callback func(models.Person)) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return StreamParseXML(file, callback)
}
