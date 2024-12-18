package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

// Person struct represents a single person
type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
}

// People struct represents a list of people
type People struct {
	XMLName xml.Name  `xml:"people"`
	People  []*Person `xml:"person"`
}

func streamingXMLHandler(w http.ResponseWriter, r *http.Request) {
	// Set the content-type header to "application/xml"
	w.Header().Set("Content-Type", "application/xml")
	// Create an encoder for the response writer
	encoder := xml.NewEncoder(w)
	// Start encoding the XML document
	/*err := encoder.EncodeToken(xml.StartElement{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}*/
	// Encode the root element (people)
	err := encoder.EncodeToken(xml.StartElement{Name: xml.Name{Local: "people"}})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Assuming you have a large dataset of people, we'll stream it in chunks
	// For demonstration purposes, we'll create a few Person instances manually
	person1 := &Person{Name: "Alice", Age: 25}
	person2 := &Person{Name: "Bob", Age: 30}
	person3 := &Person{Name: "Charlie", Age: 22}
	people := []*Person{person1, person2, person3}
	// Encode each person in the people slice
	for _, person := range people {
		err = encoder.Encode(person)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	// Encode the closing element for people
	err = encoder.EncodeToken(xml.EndElement{Name: xml.Name{Local: "people"}})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Flush the encoder to ensure all data is written to the response
	err = encoder.Flush()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/people", streamingXMLHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
