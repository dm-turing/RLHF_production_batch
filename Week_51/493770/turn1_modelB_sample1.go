package main

import (
	"fmt"
	"net/url"
	"strconv"
)

// Command interface
type Command interface {
	Execute(values url.Values) error
}

// SearchCommand struct
type SearchCommand struct {
	Query string
}

// Execute method for SearchCommand
func (c *SearchCommand) Execute(values url.Values) error {
	c.Query = values.Get("q")
	return nil
}

// FilterCommand struct
type FilterCommand struct {
	Category string
	Price    float64
}

// Execute method for FilterCommand
func (c *FilterCommand) Execute(values url.Values) error {
	c.Category = values.Get("category")
	priceStr := values.Get("price")
	if priceStr != "" {
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return err
		}
		c.Price = price
	}
	return nil
}

// CommandParser struct
type CommandParser struct {
	Commands map[string]Command
}

// NewCommandParser function
func NewCommandParser() *CommandParser {
	return &CommandParser{
		Commands: map[string]Command{
			"search": &SearchCommand{},
			"filter": &FilterCommand{},
		},
	}
}

// Parse function
func (p *CommandParser) Parse(query string) (Command, error) {
	values, err := url.ParseQuery(query)
	if err != nil {
		return nil, err
	}
	commandName := values.Get("command")
	command, ok := p.Commands[commandName]
	if !ok {
		return nil, fmt.Errorf("unknown command: %s", commandName)
	}
	err = command.Execute(values)
	if err != nil {
		return nil, err
	}
	return command, nil
}

// Example usage
func main() {
	parser := NewCommandParser()
	query := "?command=search&q=hello+world"
	command, err := parser.Parse(query)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	switch cmd := command.(type) {
	case *SearchCommand:
		fmt.Println("Search query:", cmd.Query)
	case *FilterCommand:
		fmt.Println("Filter category:", cmd.Category)
		fmt.Println("Filter price:", cmd.Price)
	}
}
