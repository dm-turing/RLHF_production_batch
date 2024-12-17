package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"sync"
)

type Person struct {
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}

type People struct {
	Persons []Person `xml:"person"`
}

func main() {
	// Simulating a large XML file
	data := `<people>
        <person>
            <name>John Doe</name>
            <age>30</age>
        </person>
        <person>
            <name>Jane Doe</name>
            <age>25</age>
        </person>
        <!-- More person entries -->
    </people>`

	var people People
	if err := xml.Unmarshal([]byte(data), &people); err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshalling XML: %v\n", err)
		os.Exit(1)
	}

	var wg sync.WaitGroup
	results := make(chan Person, len(people.Persons))
	var mu sync.Mutex
	var processed []Person

	for _, person := range people.Persons {
		wg.Add(1)
		go func(p Person) {
			defer wg.Done()
			// Simulate some processing
			processedPerson := Person{Name: p.Name, Age: p.Age + 1} // Example processing
			results <- processedPerson
		}(person)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for p := range results {
		mu.Lock()
		processed = append(processed, p)
		mu.Unlock()
	}

	// Final output
	for _, p := range processed {
		fmt.Printf("Processed Person: %+v\n", p)
	}
}
