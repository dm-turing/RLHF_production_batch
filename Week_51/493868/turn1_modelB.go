package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"compress/gzip"
)

// define the data structure for XML data
type XMLData struct {
	Field1 string `xml:"field1"`
	Field2 string `xml:"field2"`
}

// define the data structure for JSON data
type JSONData struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
}

// The TransformMiddleware function is the middleware that performs data transformation.
func TransformMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Perform data format conversion
		contentType := r.Header.Get("Content-Type")
		if strings.Contains(contentType, "application/xml") {
			transformXMLToJSON(w, r)
		} else if strings.Contains(contentType, "application/json") {
			transformJSONToGZIP(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func transformXMLToJSON(w http.ResponseWriter, r *http.Request) {
	// Read the XML request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Unmarshal XML data
	var xmlData XMLData
	err = xml.Unmarshal(body, &xmlData)
	if err != nil {
		log.Printf("Error unmarshaling XML data: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Marshal JSON data
	jsonData, err := json.Marshal(xmlData)
	if err != nil {
		log.Printf("Error marshaling JSON data: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	_, err = w.Write(jsonData)
	if err != nil {
		log.Printf("Error writing JSON response: %v", err)
	}
}

func transformJSONToGZIP(w http.ResponseWriter, r *http.Request) {
	// Read the JSON request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Compress the JSON data using gzip
	w.Header().Set("Content-Encoding", "gzip")
	gz := gzip.NewWriter(w)
	defer gz.Close()

	// Write the compressed JSON response
	_, err = gz.Write(body)
	if err != nil {
		log.Printf("Error writing gzipped JSON response: %v", err)
	}
}

func main() {
	// Example usage of the middleware
	mux := http.NewServeMux()
	mux.Handle("/", TransformMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})))

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
