package main

import (
	"fmt"
	"net/url"
)

// ParseAgeCommand implementation of the Command interface
type ParseAgeCommand struct {
	name string
}

func NewParseAgeCommand(name string) *ParseAgeCommand {
	return &ParseAgeCommand{name: name}
}

// ParseEmailCommand implementation of the Command interface
type ParseEmailCommand struct {
	name string
}

func NewParseEmailCommand(name string) *ParseEmailCommand {
	return &ParseEmailCommand{name: name}
}

// Command interface
type Command interface {
	Execute(query url.Values) error
}

// Builder interface
type Builder interface {
	AddParameter(key string, value interface{})
	GetResult() map[string]interface{}
}

// Concrete Builder implementation
type MapBuilder struct {
	result map[string]interface{}
}

func NewMapBuilder() *MapBuilder {
	return &MapBuilder{result: make(map[string]interface{})}
}

func (b *MapBuilder) AddParameter(key string, value interface{}) {
	b.result[key] = value
}

func (b *MapBuilder) GetResult() map[string]interface{} {
	return b.result
}

// Factory method for creating commands
func CreateCommand(query url.Values, parameter string) Command {
	switch parameter {
	case "age":
		return NewParseAgeCommand(query)
	case "email":
		return NewParseEmailCommand(query)
	default:
		return NewParseDefaultCommand(query, parameter)
	}
}

func main() {
	query := url.Values{"age": []string{"25"}, "email": []string{"example@example.com"}, "name": []string{"Alice"}}

	builder := NewMapBuilder()

	for param := range query {
		cmd := CreateCommand(query, param)
		if err := cmd.Execute(query); err != nil {
			fmt.Println(err)
		}
		cmd.BuildResult(builder)
	}

	result := builder.GetResult()
	fmt.Println(result) // Output: map[age:25 email:example@example.com name:Alice]
}
