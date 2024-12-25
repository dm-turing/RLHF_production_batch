package main

import (
	"fmt"
	"net/url"
)

func main() {
	// URL with query parameters
	urlString := "https://example.com/page?lang=zh&pageTitle=家"

	// Parse the URL
	u, err := url.Parse(urlString)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	// Get query values
	values := u.Query()

	// Extract i18n parameters
	lang := values.Get("lang")
	pageTitle := values.Get("pageTitle")

	// Display the results
	fmt.Printf("Language: %s, Page Title: %s \n", lang, pageTitle)
}
