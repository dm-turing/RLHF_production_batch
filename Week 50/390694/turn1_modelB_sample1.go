package main

import (
	"encoding/xml"
	"fmt"
)

type Element interface {
	String() string
}

type Book struct {
	XMLName xml.Name `xml:"book"`
	Title   string   `xml:"title"`
}

func (b Book) String() string {
	return fmt.Sprintf("Book: %s", b.Title)
}

type Movie struct {
	XMLName xml.Name `xml:"movie"`
	Title   string   `xml:"title"`
	Year    int      `xml:"year"`
}

func (m Movie) String() string {
	return fmt.Sprintf("Movie: %s (%d)", m.Title, m.Year)
}

func createElement(data []byte) (Element, error) {
	var e interface{}
	switch string(data[:4]) { // Simple heuristic, replace with proper logic
	case "<boo":
		e = &Book{}
	case "<mov":
		e = &Movie{}
	default:
		return nil, fmt.Errorf("unknown element type")
	}

	if err := xml.Unmarshal(data, e); err != nil {
		return nil, err
	}

	return e.(Element), nil
}

func main() {
	bookData := []byte(`<book><title>The Go Programming Language</title></book>`)
	movieData := []byte(`<movie><title>Inception</title><year>2010</year></movie>`)

	book, err := createElement(bookData)
	if err != nil {
		panic(err)
	}
	fmt.Println(book)

	movie, err := createElement(movieData)
	if err != nil {
		panic(err)
	}
	fmt.Println(movie)
}
