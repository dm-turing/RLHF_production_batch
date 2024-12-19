package main

import (
	"fmt"
	"net/url"
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

// Example usage:
func ExampleQueryBuilder_Build() {
	qb := New()

	qb.Set("sort", "name") // Sort the results by "name"
	qb.Add("category", "technology")
	qb.Add("category", "culture")
	qb.Set("limit", "10")

	// Generated query string: sort=name&category=technology&category=culture&limit=10
	query := qb.Build()

	fmt.Println(query)
	// Output: sort=name&category=technology&category=culture&limit=10
}

func main() {
	ExampleQueryBuilder_Build()
}
