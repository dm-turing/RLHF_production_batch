package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	"time"
)

type User struct {
	XMLName xml.Name `xml:"user"`
	Name    string   `xml:"name"`
	Email   string   `xml:"email"`
	Age     int      `xml:"age"`
}

type UserJSON struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func processJSONFile(filename string, wg *sync.WaitGroup) {
	defer wg.Done()
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	var user UserJSON
	if err := json.Unmarshal(data, &user); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
	fmt.Println("Processed JSON User:", user)
}

func processXMLFile(filename string, wg *sync.WaitGroup) {
	defer wg.Done()
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading XML file: %v", err)
	}

	var user User
	if err := xml.Unmarshal(data, &user); err != nil {
		log.Fatalf("Error decoding XML: %v", err)
	}
	fmt.Println("Processed XML User:", user)
}

func main() {
	start := time.Now()
	jsonFiles := []string{"user.json"}
	xmlFiles := []string{"user.xml"}

	var wg sync.WaitGroup
	for _, file := range jsonFiles {
		wg.Add(1)
		go processJSONFile(file, &wg)
	}

	for _, file := range xmlFiles {
		wg.Add(1)
		go processXMLFile(file, &wg)
	}

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %s\n", elapsed)
}
