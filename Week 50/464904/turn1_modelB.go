package main

import (
	"fmt"
	"net/url"
	"strconv"
)

type Query struct {
	Filter map[string]string `json:"filter"`
	Page   int               `json:"page"`
	Limit  int               `json:"limit"`
}

func NewQuery(values url.Values) *Query {
	query := &Query{
		Filter: map[string]string{},
		Page:   1,
		Limit:  10,
	}

	// Parse query parameters
	for key, value := range values {
		if len(value) == 1 {
			switch key {
			case "page":
				page, err := strconv.Atoi(value[0])
				if err == nil {
					query.Page = page
				}
			case "limit":
				limit, err := strconv.Atoi(value[0])
				if err == nil {
					query.Limit = limit
				}
			default:
				query.Filter[key] = value[0]
			}
		}
	}
	return query
}
func main() {
	// Example usage
	urlValues := url.Values{}
	urlValues.Add("page", "2")
	urlValues.Add("limit", "15")
	urlValues.Add("name", "John")
	urlValues.Add("age", "30")

	query := NewQuery(urlValues)
	fmt.Println("Query:", query)

	// Perform filtering and pagination based on the query parameters
	// ...
}
