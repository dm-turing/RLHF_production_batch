package main

import (
	"fmt"
)

// User struct
type User struct {
	Name   string   `xml:"name"`
	Groups []*Group `xml:"groups>group,omitempty"`
}

// Group struct
type Group struct {
	Name  string  `xml:"name"`
	Users []*User `xml:"users>user,omitempty"`
}

type Encoder struct {
	seen map[interface{}]bool
}

func (e *Encoder) MarshalXML(v interface{}) (string, error) {
	e.seen = make(map[interface{}]bool)
	return e.encode(v)
}

func (e *Encoder) encode(v interface{}) (string, error) {
	switch value := v.(type) {
	case *User:
		if e.seen[value] {
			return "<User>...</User>", nil // Reference already seen
		}
		e.seen[value] = true // Mark as seen
		return fmt.Sprintf("<User><name>%s</name></User>", value.Name), nil
	case *Group:
		if e.seen[value] {
			return "<Group>...</Group>", nil // Reference already seen
		}
		e.seen[value] = true // Mark as seen
		groupXML := fmt.Sprintf("<Group><name>%s</name>", value.Name)
		for _, user := range value.Users {
			userXML, _ := e.encode(user)
			groupXML += userXML
		}
		groupXML += "</Group>"
		return groupXML, nil
	}
	return "", fmt.Errorf("unsupported type")
}

func main() {
	// Create instances
	user1 := &User{Name: "Alice"}
	user2 := &User{Name: "Bob"}
	group1 := &Group{Name: "Developers", Users: []*User{user1, user2}}
	user1.Groups = []*Group{group1} // Circular reference

	encoder := &Encoder{}
	xmlData, err := encoder.MarshalXML(group1)
	if err != nil {
		fmt.Println("Error encoding to XML:", err)
		return
	}
	fmt.Println(xmlData)
}
