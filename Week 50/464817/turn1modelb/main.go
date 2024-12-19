package main

import (
	"fmt"
	"querybuilder"
)

// Example usage:
func ExampleQueryBuilder_Build() {
	qb := querybuilder.New()

	qb.Set("sort", "name") // Sort the results by "name"
	qb.Add("category", "technology")
	qb.Add("category", "culture")
	qb.Set("limit", "10")

	// Generated query string: sort=name&category=technology&category=culture&limit=10
	query := qb.Build()

	fmt.Println(query)
	// Output: sort=name&category=technology&category=culture&limit=10
}
