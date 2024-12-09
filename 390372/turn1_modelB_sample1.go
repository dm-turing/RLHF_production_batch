package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func processJSONFile(filename string, wg *sync.WaitGroup) {
	defer wg.Done()

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var person Person
	if err := json.Unmarshal(data, &person); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	fmt.Printf("Processed JSON: %+v\n", person)
}

func main() {
	dir := "./json_files" // Replace this with your JSON files directory

	filenames, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	var wg sync.WaitGroup

	for _, filename := range filenames {
		if filename.IsDir() {
			continue
		}

		filePath := filepath.Join(dir, filename.Name())
		wg.Add(1)
		go processJSONFile(filePath, &wg)
	}

	wg.Wait()
	fmt.Println("JSON processing completed.")
}
