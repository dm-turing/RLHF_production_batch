package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type RequestData struct {
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}

func main() {
	http.HandleFunc("/data", handler)
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	middleware(handlerFunc)(w, r)
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	// Assuming the request body is in XML format
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var data RequestData
	body = []byte(`<data><name>Immortal</name><age>100</age></data>`)
	if err := xml.Unmarshal(body, &data); err != nil {
		log.Println("Error unmarshaling XML:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Do something with the data
	fmt.Fprintf(w, "Name: %s, Age: %d\n", data.Name, data.Age)
}

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check the Content-Type header to decide if conversion is necessary
		contentType := r.Header.Get("Content-Type")
		if contentType == "application/xml" {
			// Convert XML body to JSON
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Println("Error reading request body:", err)
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			var data RequestData
			if err := xml.Unmarshal(body, &data); err != nil {
				log.Println("Error unmarshaling XML:", err)
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}

			jsonBody, err := json.Marshal(data)
			if err != nil {
				log.Println("Error marshaling JSON:", err)
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}

			// Modify the request to have a JSON body
			r.Body = ioutil.NopCloser(bytes.NewReader(jsonBody))
			r.Header.Set("Content-Type", "application/json")
		}

		// Proceed to the next handler
		next(w, r)
	}
}
