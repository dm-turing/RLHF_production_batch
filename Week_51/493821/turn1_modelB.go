package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	rawURL := "https://example.com?name=Carlos&age=27&language=es-ES&location=España"
	u, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}

	query := u.Query()

	// Handle non-i18n parameters
	name := query.Get("name")
	age := query.Get("age")
	fmt.Printf("Name: %s, Age: %s\n", name, age)

	// Handle i18n parameters
	language := query.Get("language")
	location := query.Get("location")

	// Option 1: Using strings.Replace
	language = strings.Replace(language, "-", "_", -1)
	location = strings.Replace(location, " ", "_", -1)

	// Option 2: Using a library (recommended)
	// import "github.com/go-playground/locales/en"
	// import "golang.org/x/text/language"
	// tag, _ := language.Parse(language)
	// language = tag.String()
	// location = tag.String()

	fmt.Printf("Language: %s, Location: %s\n", language, location)
}
