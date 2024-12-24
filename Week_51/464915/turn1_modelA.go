package main

import (
	"fmt"
	"net/url"
	"strconv"
)

// ParseQuery parse the URL query parameters to a map with their respective types.
func ParseQuery(urlStr string) (map[string]interface{}, error) {
	url, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	parsedQuery := url.Query()

	result := make(map[string]interface{})

	for key, values := range parsedQuery {
		switch {
		case len(values) == 1:
			v := values[0]
			// Check if the value is an integer or float
			if i, err := strconv.Atoi(v); err == nil {
				result[key] = i
			} else if f, err := strconv.ParseFloat(v, 64); err == nil {
				result[key] = f
			} else {
				result[key] = v
			}
		case len(values) > 1:
			sliceValue := make([]interface{}, len(values))
			for i, v := range values {
				// Check if the slice element is an integer or float
				if j, err := strconv.Atoi(v); err == nil {
					sliceValue[i] = j
				} else if f, err := strconv.ParseFloat(v, 64); err == nil {
					sliceValue[i] = f
				} else {
					sliceValue[i] = v
				}
			}
			result[key] = sliceValue
		}
	}

	return result, nil
}

func main() {
	queryURL := "http://example.com/api?id=123&name=John&price=29.99&numbers=1&numbers=2&numbers=3"
	parsedQuery, err := ParseQuery(queryURL)
	if err != nil {
		fmt.Println("Error parsing query:", err)
		return
	}

	fmt.Println("Parsed Query:")
	for key, value := range parsedQuery {
		fmt.Printf("%s: %v\n", key, value)
	}
}
