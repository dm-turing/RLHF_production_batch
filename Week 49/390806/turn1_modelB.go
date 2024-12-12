package main

import (
	"encoding/xml"
	"fmt"
	"reflect"
)

func reflectToXML(v interface{}) ([]byte, error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Struct {
		return nil, fmt.Errorf("unsupported type: %T", v)
	}

	result := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<root>\n"

	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		fieldType := rv.Type().Field(i)

		// Check if field should be omitted
		if field.IsZero() && fieldType.Tag.Get("omitempty") == "true" {
			continue
		}

		// Get the XML tag name
		tagName := fieldType.Tag.Get("xml")
		if tagName == "" {
			tagName = fieldType.Name
		}

		// Encode field value
		switch field.Kind() {
		case reflect.Struct:
			subXML, err := reflectToXML(field.Interface())
			if err != nil {
				return nil, err
			}
			result += fmt.Sprintf("\t<%s>%s</%s>\n", tagName, string(subXML), tagName)
		default:
			result += fmt.Sprintf("\t<%s>%v</%s>\n", tagName, field.Interface(), tagName)
		}
	}
	result += "</root>"
	return []byte(result), nil
}

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Address struct {
		Street string `xml:"street"`
		City   string `xml:"city"`
	} `xml:"address"`
}

func main() {
	p := Person{
		Name: "John Doe",
		Age:  30,
		Address: struct {
			Street string `xml:"street"`
			City   string `xml:"city"`
		}{
			Street: "Main St",
			City:   "Anytown",
		},
	}

	xmlBytes, err := reflectToXML(p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(xmlBytes))
}
