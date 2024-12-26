package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User1 struct {
	Name string
	Age  int
}

func ServeHTML(w http.ResponseWriter, r *http.Request) {

	t, _ := template.New("t").Parse("<html><head><title>User Profile</title></head><body>    <h1>Hello, {{.Name}}!</h1>    <p>You are {{.Age}} years old.</p></body></html>")
	tmpl := template.Must(t, fmt.Errorf("%s", "Error parsing the template"))
	user := User1{
		Name: "Bob", // Replace with actual user input
		Age:  25,    // Replace with actual user input
	}

	if err := tmpl.Execute(w, user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", ServeHTML)
	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", http.DefaultServeMux)
}
