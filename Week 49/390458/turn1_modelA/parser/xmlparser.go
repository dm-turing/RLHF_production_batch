package parser

import (
	"encoding/xml"
	"io"
	"os"

	"github.com/example/models" // Make sure to update module import paths
)

func ParseXML(data io.Reader) (*models.People, error) {
	var people models.People
	decoder := xml.NewDecoder(data)

	err := decoder.Decode(&people)
	if err != nil {
		return nil, err
	}

	return &people, nil
}

func ParseFile(filename string) (*models.People, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return ParseXML(file)
}
