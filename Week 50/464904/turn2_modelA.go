package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
)

type Query struct {
	Filter     map[string]string `json:"filter"`
	Aggregates map[string]string `json:"aggregates"`
	Page       int               `json:"page"`
	Limit      int               `json:"limit"`
}

func NewQuery(values url.Values) *Query {
	query := &Query{
		Filter:     map[string]string{},
		Aggregates: map[string]string{},
		Page:       1,
		Limit:      10,
	}

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
				if idx := strings.Index(key, ":"); idx != -1 {
					attribute, operator := key[:idx], key[idx+1:]
					query.Filter[attribute] = operator + ":" + value[0]
				} else {
					query.Aggregates[key] = value[0]
				}
			}
		}
	}

	return query
}

func main() {
	urlValues := url.Values{}
	urlValues.Add("page", "2")
	urlValues.Add("limit", "15")
	urlValues.Add("name", "John")
	urlValues.Add("age", "gt:30")
	urlValues.Add("salary", "sum")
	urlValues.Add("age", "max")

	query, err := json.MarshalIndent(NewQuery(urlValues), "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(query))
}
