package main

import (
	"fmt"
	"os"
	"text/template"
	"time"
)

func main() {
	tpl := `Today is {{formatDate .}}.`

	type Date struct {
		Time time.Time
	}

	// date := Date{time.Now()}

	formatDate := func(t time.Time) string {
		return t.Format("January 2, 2006")
	}

	tmpl := template.Must(template.New("date").Funcs(template.FuncMap{"formatDate": formatDate}).Parse(tpl))

	err := tmpl.Execute(os.Stdout, time.Now())
	if err != nil {
		fmt.Println("Error executing template:", err)
	}
}
