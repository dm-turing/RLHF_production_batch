package main

import (
	"fmt"
	"net/url"
	"strconv"
)

// Command interface
type Command interface {
	Execute(query url.Values)
}

// ParsingCommand struct implements Command
type ParsingCommand struct {
	strategy Strategy
}

func NewParsingCommand(strategy Strategy) *ParsingCommand {
	return &ParsingCommand{strategy: strategy}
}

func (c *ParsingCommand) Execute(query url.Values) {
	result := c.strategy.Parse(query)
	// Use the result as needed
	fmt.Println("result: ", result)
}

// Strategy interface
type Strategy interface {
	Parse(query url.Values) map[string]string
}

// BasicStrategy struct implements Strategy
type BasicStrategy struct{}

func (s *BasicStrategy) Parse(query url.Values) map[string]string {
	parsedData := make(map[string]string)
	for key, value := range query {
		parsedData[key] = value[0]
	}
	return parsedData
}

// Parameter interface
type Parameter interface {
	SetValue(value string)
	GetValue() string
}

// StringParameter struct implements Parameter
type StringParameter struct {
	value string
}

func (s *StringParameter) SetValue(value string) {
	s.value = value
}

func (s *StringParameter) GetValue() string {
	return s.value
}

// IntegerParameter struct implements Parameter
type IntegerParameter struct {
	value int
}

func (s *IntegerParameter) SetValue(value string) {
	v, _ := strconv.Atoi(value)
	s.value = v
}

func (s *IntegerParameter) GetValue() string {
	return strconv.Itoa(s.value)
}

// ParameterFactory method
func CreateParameter(typeName string) Parameter {
	switch typeName {
	case "string":
		return &StringParameter{}
	case "integer":
		return &IntegerParameter{}
	default:
		panic("Unsupported parameter type")
	}
}

// QueryBuilder struct
type QueryBuilder struct {
	parameters map[string]Parameter
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{parameters: make(map[string]Parameter)}
}

func (b *QueryBuilder) AddParameter(key, typeName string, value string) {
	param := CreateParameter(typeName)
	param.SetValue(value)
	b.parameters[key] = param
}

func (b *QueryBuilder) Build() map[string]string {
	result := make(map[string]string)
	for key, param := range b.parameters {
		result[key] = param.GetValue()
	}
	return result
}

func main() {
	query := url.Values{"age": []string{"25"}, "email": []string{"example@example.com"}}

	strategy := &BasicStrategy{}
	command := NewParsingCommand(strategy)

	builder := NewQueryBuilder()
	builder.AddParameter("age", "integer", query["age"][0])
	builder.AddParameter("email", "string", query["email"][0])

	parsedQuery := builder.Build()

	command.Execute(query)
	// Print parsedQuery as needed
	fmt.Println(parsedQuery)
}
