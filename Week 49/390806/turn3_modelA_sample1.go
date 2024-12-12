package main

import (
	"fmt"
	"html"
	"strings"
)

// Person struct to be encoded to XML
type Person struct {
	Name  string `xml:"name"`
	Age   int    `xml:"age"`
	Email string `xml:"email"`
}

// EncodePerson encodes a Person struct to XML without using reflection
func EncodePerson(p Person) (string, error) {
	var sb strings.Builder
	sb.WriteString("<Person>")
	sb.WriteString(fmt.Sprintf("<name>%s</name>", xmlEscape(p.Name)))
	sb.WriteString(fmt.Sprintf("<age>%d</age>", p.Age))
	sb.WriteString(fmt.Sprintf("<email>%s</email>", xmlEscape(p.Email)))
	sb.WriteString("</Person>")
	return sb.String(), nil
}

// xmlEscape escapes special XML characters in a string
func xmlEscape(s string) string {
	return html.EscapeString(s)
}

func main() {
	person := Person{Name: "John Doe", Age: 30, Email: "john@example.com"}
	xmlData, err := EncodePerson(person)
	if err != nil {
		fmt.Println("Error encoding to XML:", err)
		return
	}
	fmt.Println(xmlData)
}
