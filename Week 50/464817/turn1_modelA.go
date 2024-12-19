package main

import (
	"fmt"
	"net/url"
)

// QueryBuilder is the main struct that builds the query.
type QueryBuilder struct {
	params url.Values
}

// NewQueryBuilder returns a new QueryBuilder instance.
func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		params: url.Values{},
	}
}

// AddParam adds a new parameter to the query.
func (qb *QueryBuilder) AddParam(key string, value interface{}) *QueryBuilder {
	qb.params.Set(key, fmt.Sprint(value))
	return qb
}

// Build constructs and returns the URL query string.
func (qb *QueryBuilder) Build() string {
	return qb.params.Encode()
}

func main() {
	qb := NewQueryBuilder()
	qb.AddParam("sort", "desc")
	qb.AddParam("limit", 10)
	qb.AddParam("offset", 0)

	queryString := qb.Build()
	fmt.Println("Query String:", queryString)
}
