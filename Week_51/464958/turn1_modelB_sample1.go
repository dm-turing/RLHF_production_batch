package main

import (
	"fmt"
	"html"
	"net/url"
)

func main() {
	query := url.Values{}
	query.Add("name", "<script>alert('XSS')</script>")
	fmt.Println(html.EscapeString(query.Get("name")))
}
