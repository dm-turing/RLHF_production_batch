package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	// "regexp"
	"strings"
)

type User struct {
	Name string
}

// Custom template syntax:
// {{ variable }} for inserting a variable
// {{# loop }}...{{/ loop }} for loops
// {{? condition }}...{{?}} for conditionals

// RenderTemplate uses fmt.Sprintf to render a template
func RenderTemplate(template string, data interface{}) string {
	// Escape HTML characters
	var escapedTemplate strings.Builder
	for _, char := range template {
		switch char {
		case '<', '>', '&', '"', '\'':
			escapedTemplate.WriteString(html.EscapeString(string(char)))
		default:
			escapedTemplate.WriteRune(char)
		}
	}

	// Simple placeholder replacement
	return fmt.Sprintf(escapedTemplate.String(), data)
}

func handler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: r.URL.Query().Get("name")}
	template := `
		<html>
			<head>
				<title>Welcome</title>
			</head>
			<body>
				<h1>Hello, {{ Name }}!</h1>
				<p>Welcome to our website.</p>
			</body>
		</html>
	`
	renderedContent := RenderTemplate(template, user)
	w.Write([]byte(renderedContent))
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
