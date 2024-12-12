package main

import (
	"encoding/xml"
	"fmt"
)

// User struct
type User struct {
	Name   string   `xml:"name"`
	Groups []*Group `xml:"groups>group"`
}

// Group struct
type Group struct {
	Name  string  `xml:"name"`
	Users []*User `xml:"users>user"`
}

func main() {
	// Create instances
	user1 := &User{Name: "Alice"}
	user2 := &User{Name: "Bob"}
	group1 := &Group{Name: "Developers", Users: []*User{user1, user2}}
	user1.Groups = []*Group{group1} // Circular reference

	// Try to encode to XML
	xmlData, err := xml.MarshalIndent(group1, "", "  ")
	if err != nil {
		fmt.Println("Error encoding to XML:", err)
		return
	}
	fmt.Println(string(xmlData))
}
