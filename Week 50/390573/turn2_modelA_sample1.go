package main

import (
	"encoding/xml"
	"log"
	"net/http"
)

type Item struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type Items struct {
	XMLName xml.Name `xml:"items"`
	Items   []Item   `xml:"item"`
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	// Example data to encode
	items := []Item{
		{ID: 1, Name: "Apple"},
		{ID: 2, Name: "Banana"},
		{ID: 3, Name: "Orange"},
	}

	// Create an XML encoder
	enc := xml.NewEncoder(w)
	defer enc.Close()

	// Write the XML declaration
	// err := enc.EncodeToken(xml.StartElement{Name: xml.Name{"Version": "EncodingUTF8"}})
	// if err != nil {
	// log.Fatal(err)
	// }

	// Write the root element
	err := enc.EncodeToken(xml.StartElement{Name: xml.Name{Local: "items"}})
	if err != nil {
		log.Fatal(err)
	}

	// Write each item in a loop to avoid memory overload
	for _, item := range items {
		err = enc.Encode(item)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Write the end of the root element
	err = enc.EncodeToken(xml.EndElement{Name: xml.Name{Local: "items"}})
	if err != nil {
		log.Fatal(err)
	}

	// Write the end of XML declaration
	// err = enc.EncodeToken(xml.EndElement{})
	// err = enc.EncodeToken(xml.EndDecl())
	// if err != nil {
	// log.Fatal(err)
	// }
}

func main() {
	http.HandleFunc("/items", itemsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
