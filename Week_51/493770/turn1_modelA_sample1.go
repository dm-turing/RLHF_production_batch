package main

import (
	"fmt"
	"net/url"
)

// Command interface
type Command interface {
	Execute(query string) (map[string]string, error)
}

// ParseCommand struct implements Command
type ParseCommand struct {
	urlParser URLParser
}

func NewParseCommand(parser URLParser) *ParseCommand {
	return &ParseCommand{urlParser: parser}
}

func (c *ParseCommand) Execute(query string) (map[string]string, error) {
	return c.urlParser.Parse(query)
}

// URLParser interface
type URLParser interface {
	Parse(query string) (map[string]string, error)
}

// DefaultURLParser struct implements URLParser
type DefaultURLParser struct {
}

func (p *DefaultURLParser) Parse(query string) (map[string]string, error) {
	u, err := url.ParseQuery(query)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func main() {

	parser := NewParseCommand(DefaultURLParser{})
	result, err := parser.Execute("name=John&age=30&city=NewYork")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed Query:", result)
	}
}
