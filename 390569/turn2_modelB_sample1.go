package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	tpl := `Hello, {{.Name}}! Your age is {{.Age}}.`

	type Person struct {
		Name string
		Age  int
	}

	person := Person{"Alice", 30}

	tmpl := template.Must(template.New("person").Parse(tpl))

	err := tmpl.Execute(os.Stdout, person)
	if err != nil {
		fmt.Println("Error executing template:", err)
	}
}
