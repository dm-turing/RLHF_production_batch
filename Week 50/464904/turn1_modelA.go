package main

import (
	"fmt"
	"net/url"
)

// QueryBuilder is a struct to hold query parameters
type QueryBuilder struct {
	offset  *int
	limit   *int
	filters map[string]string
	sort    string
}

// NewQueryBuilder initializes a new QueryBuilder instance
func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		offset:  nil,
		limit:   nil,
		filters: make(map[string]string),
		sort:    "",
	}
}

// WithOffset sets the offset for pagination
func (qb *QueryBuilder) WithOffset(offset int) *QueryBuilder {
	qb.offset = &offset
	return qb
}

// WithLimit sets the limit for pagination
func (qb *QueryBuilder) WithLimit(limit int) *QueryBuilder {
	qb.limit = &limit
	return qb
}

// WithFilter adds a key-value filter
func (qb *QueryBuilder) WithFilter(key, value string) *QueryBuilder {
	qb.filters[key] = value
	return qb
}

// WithSort sets the sorting parameter
func (qb *QueryBuilder) WithSort(sort string) *QueryBuilder {
	qb.sort = sort
	return qb
}

// Build constructs the query string based on the builder settings
func (qb *QueryBuilder) Build() string {
	v := url.Values{}

	if qb.offset != nil {
		v.Set("offset", fmt.Sprint(*qb.offset))
	}

	if qb.limit != nil {
		v.Set("limit", fmt.Sprint(*qb.limit))
	}

	for key, value := range qb.filters {
		v.Set(key, value)
	}

	if qb.sort != "" {
		v.Set("sort", qb.sort)
	}

	return v.Encode()
}

func main() {
	builder := NewQueryBuilder()
	queryString := builder.
		WithOffset(0).
		WithLimit(10).
		WithFilter("name", "John").
		WithFilter("age", "30").
		WithSort("-created_at").
		Build()

	fmt.Println("Query String:", queryString)
}
