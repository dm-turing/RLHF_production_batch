package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// Attachment struct represents a single attachment
type Attachment struct {
	User      string
	Content   []byte // Storing the file content as bytes
	Timestamp time.Time
	Filetype  string
	Filesize  int64
	Geoloc    string
}

// BlogPost struct represents a blog post with its attachments
type BlogPost struct {
	Title       string
	Content     string
	Attachments []Attachment
}

func main() {
	// Create a slice to store blog posts
	var blogPosts []BlogPost

	// Create a new blog post and add attachments
	newPost := BlogPost{
		Title:   "My First Blog Post",
		Content: "Hello, this is my first blog post!",
	}

	attachment1 := Attachment{
		User:      "John Doe",
		Timestamp: time.Now(),
		Filetype:  "image/png",
	}

	// Read the file content from the disk
	fileContent, err := ioutil.ReadFile("image.png")
	if err != nil {
		log.Fatal(err)
	}
	attachment1.Content = fileContent

	// Calculate the file size
	fileInfo, err := os.Stat("image.png")
	if err != nil {
		log.Fatal(err)
	}
	attachment1.Filesize = fileInfo.Size()

	// Add the attachment to the blog post
	newPost.Attachments = append(newPost.Attachments, attachment1)

	// Add the blog post to the slice
	blogPosts = append(blogPosts, newPost)

	// Accessing and displaying the attachment details
	for _, post := range blogPosts {
		fmt.Println("Blog Post Title:", post.Title)
		fmt.Println("Blog Post Content:", post.Content)
		fmt.Println("Attachments:")
		for _, attachment := range post.Attachments {
			fmt.Println("  User:", attachment.User)
			fmt.Println("  Timestamp:", attachment.Timestamp)
			fmt.Println("  Filetype:", attachment.Filetype)
			fmt.Println("  Filesize:", attachment.Filesize)
			// To handle the attachment content, you can use various libraries
			// depending on the file type (e.g., image/png, image/jpeg, video/mp4).
		}
		fmt.Println()
	}
}
