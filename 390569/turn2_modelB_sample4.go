package main

import (
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/Masterminds/sprig"
)

func main() {
	tpl := `Today is {{ .Time | date "January 2, 2006" }}.`

	type Date struct {
		Time time.Time
	}

	date := Date{time.Now()}

	// Create a new template with Sprig functions
	tmpl := sprig.TxtFuncMap(sprig.Default)
	tmpl = template.Must(template.New("date").Funcs(tmpl).Parse(tpl))

	err := tmpl.Execute(os.Stdout, date)
	if err != nil {
		fmt.Println("Error executing template:", err)
	}
}
