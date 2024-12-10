package main

import (
	"fmt"
	"html/template"
	"os"

	"github.com/Masterminds/sprig"
)

func main() {
	tpl := `Hello, {{ .Name | capitalize }}! Your age is {{ .Age }}.`

	type Person struct {
		Name string
		Age  int
	}

	person := Person{"alice", 30}

	// Create a new template with Sprig functions
	tmpl := sprig.TxtFuncMap(sprig.Default)
	tmpl = template.Must(template.New("person").Funcs(tmpl).Parse(tpl))

	err := tmpl.Execute(os.Stdout, person)
	if err != nil {
		fmt.Println("Error executing template:", err)
	}
}
