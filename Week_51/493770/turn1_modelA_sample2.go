package main

import (
	"fmt"
	"net/url"
)

// URLParserStrategy interface
type URLParserStrategy interface {
	Parse(query string) (map[string]string, error)
}

// DefaultURLParserStrategy struct implements URLParserStrategy
type DefaultURLParserStrategy struct {
}

func (p *DefaultURLParserStrategy) Parse(query string) (map[string]string, error) {
	u, err := url.ParseQuery(query)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// CustomURLParserStrategy struct implements URLParserStrategy
type CustomURLParserStrategy struct {
}

func (p *CustomURLParserStrategy) Parse(query string) (map[string]string, error) {
	return nil, fmt.Errorf("Custom parser not implemented")
}

// URLParserContext struct
type URLParserContext struct {
	strategy URLParserStrategy
}

func NewURLParserContext(strategy URLParserStrategy) *URLParserContext {
	return &URLParserContext{strategy: strategy}
}

func (c *URLParserContext) SetStrategy(strategy URLParserStrategy) {
	c.strategy = strategy
}

func (c *URLParserContext) Parse(query string) (map[string]string, error) {
	return c.strategy.Parse(query)
}

func main() {
	context := NewURLParserContext(DefaultURLParserStrategy{})
	result, err := context.Parse("name=John&age=30&city=NewYork")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed Query:", result)
	}

	// Switch to a custom strategy (if implemented)
	context.SetStrategy(CustomURLParserStrategy{})
}
