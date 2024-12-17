package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"sync"
)

const chunkSize = 4096 // Size of the chunk to read from the XML file

type Person struct {
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}

type People struct {
	XMLName xml.Name `xml:"people"`
	Persons []Person `xml:"person"`
}

func readXMLInChunks(r io.Reader) ([]byte, error) {
	var buf bytes.Buffer
	var wg sync.WaitGroup
	var err error

	done := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				return
			default:
				data := make([]byte, chunkSize)
				n, e := r.Read(data)
				if e != nil {
					err = e
					close(done)
					return
				}
				if n == 0 {
					close(done)
					return
				}
				buf.Write(data[:n])
			}
		}
	}()

	wg.Wait()
	return buf.Bytes(), err
}

func parsePeopleConcurrently(data []byte, ch chan<- Person) {
	defer close(ch)

	// var p People
	xmlDecoder := xml.NewDecoder(bytes.NewReader(data))
	for {
		t, err := xmlDecoder.Token()
		if err != nil {
			break
		}

		if se, ok := t.(xml.StartElement); ok && se.Name.Local == "person" {
			var person Person
			if err := xmlDecoder.DecodeElement(&person, &se); err != nil {
				fmt.Println("Error decoding person:", err)
				continue
			}
			ch <- person
		}
	}
}

func main() {
	file, err := os.Open("large.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	data, err := readXMLInChunks(file)
	if err != nil {
		fmt.Println("Error reading XML in chunks:", err)
		return
	}

	var wg sync.WaitGroup
	ch := make(chan Person)

	wg.Add(1)
	go parsePeopleConcurrently(data, ch)

	go func() {
		wg.Done()
		for person := range ch {
			fmt.Printf("%+v\n", person)
		}
	}()

	wg.Wait()
}
