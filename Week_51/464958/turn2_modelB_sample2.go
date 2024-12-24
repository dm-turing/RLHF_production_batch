package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Get the user input from the query parameter
	userInput := r.URL.Query().Get("input")

	// Create a template and parse it
	t := template.Must(template.New("index").Parse(`
	<html>
	<head>
		<title>XSS Example</title>
	</head>
	<body>
		<h1>User Input:</h1>
		<p>{{ .UserInput }}</p>
	</body>
	</html>
	`))