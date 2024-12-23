package main

import (
	"fmt"
	"net/url"
)

func main() {
	query := url.Values{}
	query.Add("name", "John Doe")
	query.Add("address", "New York City")
	encodedQuery := query.Encode()
	fmt.Println(encodedQuery)
}
