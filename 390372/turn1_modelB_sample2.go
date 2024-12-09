package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"
)

type Book struct {
	XMLName xml.Name `xml:"book"`
	Title   string   `xml:"title"`
	Author  string   `xml:"author"`
}

func processXMLFile(filename string, wg *sync.WaitGroup) {
	defer wg.Done()

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var book Book
	if err := xml.Unmarshal(data, &book); err != nil {
		log.Fatalf("Error unmarshalling XML: %v", err)
	}

	fmt.Printf("Processed XML: %+v\n", book)
}

func main() {
	dir := "./xml_files" // Replace this with your XML files directory

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
		go processXMLFile(filePath, &wg)
	}

	wg.Wait()
	fmt.Println("XML processing completed.")
}
