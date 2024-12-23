package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl := template.Must(template.New("example").Parse(`
    <html>
    <body>
    <h1>Hello, my name is {{ .Name }}.</h1>
    <p>I am {{ .Age }} years old and have been programming in {{ .Language }} for {{ .Years }} years.</p>
    </body>
    </html>
    `))

	data := struct {
		Name     string
		Age      int
		Language string
		Years    int
	}{
		"Bob",
		25,
		"Java",
		3,
	}

	if err := tpl.Execute(os.Stdout, data); err != nil {
		log.Fatalf("error executing template: %v", err)
	}
}
