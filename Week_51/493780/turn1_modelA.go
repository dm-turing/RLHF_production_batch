package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type API struct {
	Version string   `json:"version"`
	Changes []string `json:"changes"`
}

func main() {
	api := &API{
		Version: "1.0.1",
		Changes: []string{
			"- Add new endpoint: /api/v1/new-endpoint",
			"- Enhance authentication logic",
			"- Fix bug #123: Error handling in /api/v1/existing-endpoint",
		},
	}

	// Generate changelog.md
	changelog, err := json.MarshalIndent(api, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	changelogString := string(changelog)

	// Write changelog to file
	err = ioutil.WriteFile("changelog.md", []byte(changelogString), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Changelog generated successfully!")
}
