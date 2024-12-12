package main

import (
	"encoding/xml"
	"fmt"
	"reflect"
)

type Address struct {
	Street string `xml:"street"`
	City   string `xml:"city"`
}

type Person struct {
	Name  string `xml:"name"`
	Age   int    `xml:"age"`
	Email string `xml:"email"`
	Address
}

// Custom MarshalXML method to minimize reflection
func (p *Person) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// Directly encode known fields
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	if err := e.EncodeElement(&p.Name, xml.StartElement{Name: xml.Name{Local: "name"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(&p.Age, xml.StartElement{Name: xml.Name{Local: "age"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(&p.Email, xml.StartElement{Name: xml.Name{Local: "email"}}); err != nil {
		return err
	}

	// Use reflection for dynamic fields in Address
	addressValue := reflect.ValueOf(p.Address)
	addressType := addressValue.Type()

	for i := 0; i < addressType.NumField(); i++ {
		field := addressValue.Field(i)
		fieldType := addressType.Field(i)

		xmlTag := fieldType.Tag.Get("xml")
		if xmlTag == "" {
			xmlTag = fieldType.Name
		}

		if err := e.EncodeElement(field.Interface(), xml.StartElement{Name: xml.Name{Local: xmlTag}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func main() {
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			Street: "123 Main St",
			City:   "Anytown",
		},
	}

	data, err := xml.Marshal(&person)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
