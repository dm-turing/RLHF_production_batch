package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type Change struct {
	Type        string
	Description string
}

type Changelog struct {
	Version string
	Changes []Change
}

func main() {
	// Define the changes for the latest API version
	changes := []Change{
		{
			Type:        "New Feature",
			Description: "Added a new endpoint to retrieve user profiles.",
		},
		{
			Type:        "Improvement",
			Description: "Enhanced the error handling for invalid requests.",
		},
		{
			Type:        "Bug Fix",
			Description: "Fixed an issue with pagination on the list endpoints.",
		},
	}

	// Create a new Changelog struct
	changelog := Changelog{
		Version: "v1.2.0",
		Changes: changes,
	}

	// Render the changelog template
	tmpl, err := template.ParseFiles("changelog.tmpl")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	f, err := os.Create("CHANGELOG.md")
	if err != nil {
		log.Fatalf("Error creating changelog file: %v", err)
	}

	defer f.Close()

	if err := tmpl.Execute(f, changelog); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	fmt.Println("Changelog generated successfully!")
}
