// Go code using text/template
package main

import (
	"os"
	"text/template"
)

func main() {
	data := struct {
		Name   string
		Age    int
		Skills []string
	}{"Alice", 30, []string{"Python", "Go", "JavaScript"}}

	tmpl, err := template.New("person").Parse(`
    Name: {{.Name}}
    Age: {{.Age}}
    Skills: {{range .Skills}}- {{.}}{{end}}
    `)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
