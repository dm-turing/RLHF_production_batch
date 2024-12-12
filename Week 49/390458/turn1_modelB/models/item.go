package models

import "encoding/xml"

// Item represents an RSS item
type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}

// RSS represents the root of an RSS feed
type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel struct {
		Title       string  `xml:"title"`
		Link        string  `xml:"link"`
		Description string  `xml:"description"`
		Items       []*Item `xml:"item"`
	} `xml:"channel"`
}
