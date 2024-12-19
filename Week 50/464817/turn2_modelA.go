package main

import (
	"fmt"
	"net/url"
	"strings"
)

// QueryBuilder provides methods to construct database queries with filtering options.
type QueryBuilder struct {
	values url.Values
}

// New creates a new QueryBuilder instance.
func New() *QueryBuilder {
	return &QueryBuilder{values: make(url.Values)}
}

// Set sets a query parameter with the given value.
func (qb *QueryBuilder) Set(param, value string) *QueryBuilder {
	qb.values.Set(param, value)
	return qb
}

// Add adds a query parameter with the given value. Multiple values can be specified.
func (qb *QueryBuilder) Add(param, value string) *QueryBuilder {
	qb.values.Add(param, value)
	return qb
}

// Del removes all query parameters with the given key.
func (qb *QueryBuilder) Del(param string) *QueryBuilder {
	qb.values.Del(param)
	return qb
}

// Build constructs the final query string.
func (qb *QueryBuilder) Build() string {
	return qb.values.Encode()
}

// QueryFilter defines a filter for a query with a specific operator and value.
type QueryFilter struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

// ApplyFilter applies a filter to the query builder.
func (qb *QueryBuilder) ApplyFilter(filter QueryFilter) *QueryBuilder {
	value := filter.Value
	if s, ok := value.(string); ok {
		value = s
	}

	switch filter.Operator {
	case "=":
		qb.Set(filter.Field, fmt.Sprintf("%v", value))
	case "!=":
		qb.Set(filter.Field, fmt.Sprintf("%v", value))
	case ">":
		qb.Set(filter.Field, fmt.Sprintf("%v", value))
	case "<":
		qb.Set(filter.Field, fmt.Sprintf("%v", value))
	case ">=":
		qb.Set(filter.Field, fmt.Sprintf("%v", value))
	case "<=":
		qb.Set(filter.Field, fmt.Sprintf("%v", value))
	case "LIKE":
		qb.Set(filter.Field, fmt.Sprintf("%v", value))
	case "IN":
		if s, ok := value.([]string); ok {
			qb.Add(filter.Field, strings.Join(s, ","))
		}
	default:
		panic(fmt.Sprintf("Invalid operator: %v", filter.Operator))
	}

	return qb
}

// Example usage:
func ExampleQueryBuilder_BuildWithFilters() {
	qb := New()

	// Apply filters
	filter1 := QueryFilter{Field: "name", Operator: "LIKE", Value: "%John%"}
	filter2 := QueryFilter{Field: "age", Operator: ">=", Value: 30}
	filter3 := QueryFilter{Field: "category", Operator: "IN", Value: []string{"technology", "culture"}}

	qb.ApplyFilter(filter1)
	qb.ApplyFilter(filter2)
	qb.ApplyFilter(filter3)

	// Generated query string: name=%John%&age=30&category=technology%2Cculture
	query := qb.Build()

	fmt.Println(query)
	// Output: name=%John%&age=30&category=technology%2Cculture
}

func main() {
	ExampleQueryBuilder_BuildWithFilters()
}
