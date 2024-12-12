package main

import (
	"encoding/xml"
	"fmt"
)

type Post struct {
	Title   string `xml:"title"`
	Content string `xml:"content"`
	Blog    *Blog  `xml:"-"` // Omit Blog in XML to avoid circular reference
}

type Blog struct {
	Name        string  `xml:"name"`
	Description string  `xml:"description"`
	Posts       []*Post `xml:"post"`
}

// Custom MarshalXML method to handle circular reference
func (b *Blog) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	if err := e.EncodeElement(&b.Name, xml.StartElement{Name: xml.Name{Local: "name"}}); err != nil {
		return err
	}

	if err := e.EncodeElement(&b.Description, xml.StartElement{Name: xml.Name{Local: "description"}}); err != nil {
		return err
	}

	for _, post := range b.Posts {
		// Temporarily set Blog to nil to avoid encoding the circular reference
		tempBlog := post.Blog
		post.Blog = nil
		if err := e.Encode(post); err != nil {
			return err
		}
		post.Blog = tempBlog
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func main() {
	post1 := &Post{Title: "Post 1", Content: "Content of Post 1"}
	post2 := &Post{Title: "Post 2", Content: "Content of Post 2"}

	blog := &Blog{
		Name:        "My Blog",
		Description: "A blog about everything.",
		Posts:       []*Post{post1, post2},
	}

	post1.Blog = blog
	post2.Blog = blog

	data, err := xml.Marshal(blog)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
