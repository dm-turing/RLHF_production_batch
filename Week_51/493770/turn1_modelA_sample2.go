package main

import (
	"fmt"
	"net/url"
)

// Strategy interface
type Strategy interface {
	Parse(query url.Values) map[string]string
}

// BasicParser implementation of the Strategy interface
type BasicParser struct{}

func (p *BasicParser) Parse(query url.Values) map[string]string {
	parsedData := make(map[string]string)
	for key, value := range query {
		parsedData[key] = value[0]
	}
	return parsedData
}

// EnhancedParser implementation of the Strategy interface
type EnhancedParser struct{}

func (p *EnhancedParser) Parse(query url.Values) map[string]string {
	parsedData := make(map[string]string)
	for key, value := range query {
		parsedData[key] = value[0]
		// Enhanced logic can go here
	}
	return parsedData
}

func main() {
	query := url.Values{"age": []string{"25"}, "email": []string{"example@example.com"}}

	strategies := []Strategy{
		&BasicParser{},
		&EnhancedParser{},
	}

	for _, strategy := range strategies {
		data := strategy.Parse(query)
		fmt.Println(data)
	}
}
