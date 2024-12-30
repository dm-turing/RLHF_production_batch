package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Use a slice of strings for simple string interpolation
	names := []string{"Alice", "Bob", "Charlie"}
	fmt.Println("Hello, ", strings.Join(names, ", "), "!") // Output: Hello, Alice, Bob, Charlie!
	// Use a slice of structs for binding dynamic data to HTML templates
	people := []Person{{"Alice", 25}, {"Bob", 30}, {"Charlie", 22}}
	tmpl := template.Must(template.New("").Parse("<html><body><h1>People</h1><ul>{{range .}}<li>{{.Name}} - {{.Age}}</li>{{end}}</ul></body></html>"))
	tmpl.Execute(os.Stdout, people) // Output: <html><body><h1>People</h1><ul><li>Alice - 25</li><li>Bob - 30</li><li>Charlie - 22</li></ul></body></html>

}
