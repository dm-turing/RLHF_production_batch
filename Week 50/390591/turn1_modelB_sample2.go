package main

import (
	"encoding/xml"
	"fmt"
	"regexp"
)

type User struct {
	XMLName xml.Name `xml:"user"`
	Name    string   `xml:"name"`
	Email   string   `xml:"email" validate:"email"`
}

func main() {
	// Example XML data
	xmlData := []byte(`<user><name>Joe</name><email>joe@example.com</email></user>`)
	var user User
	if err := xml.Unmarshal(xmlData, &user); err != nil {
		fmt.Println("Error unmarshaling XML:", err)
		return
	}
	// Validate email format using regex after unmarshaling
	if err := validateEmail(user.Email); err != nil {
		fmt.Println("Invalid email format:", err)
	} else {
		fmt.Println("Email format is valid")
	}
}

func validateEmail(email string) error {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if !regexp.MustCompile(emailRegex).MatchString(email) {
		return fmt.Errorf("Invalid email format: %s", email)
	}
	return nil
}
