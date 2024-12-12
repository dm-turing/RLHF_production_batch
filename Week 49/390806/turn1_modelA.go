package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Person struct {
	Name  string `xml:"name"`
	Age   int    `xml:"age"`
	Email string `xml:"email"`
}

func EncodeToXML(v interface{}) (string, error) {
	value := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	if typ.Kind() != reflect.Struct {
		return "", fmt.Errorf("expected a struct, got %s", typ.Kind())
	}

	elements := make([]string, 0)
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldType := typ.Field(i)

		xmlTag := fieldType.Tag.Get("xml")
		if xmlTag == "" {
			xmlTag = fieldType.Name
		}

		elements = append(elements, fmt.Sprintf("<%s>%v</%s>", xmlTag, field.Interface(), xmlTag))
	}
	return fmt.Sprintf("<Person>%s</Person>", strings.Join(elements, "")), nil
}

func main() {
	person := Person{Name: "John Doe", Age: 30, Email: "john@example.com"}
	xmlData, err := EncodeToXML(person)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(xmlData)
}
