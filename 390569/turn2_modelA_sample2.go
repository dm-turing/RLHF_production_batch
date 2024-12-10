package main

import (
	"html/template"
	"os"
	"time"

	"github.com/Masterminds/sprig/v3"
)

// Data structure to hold the report data
type Report struct {
	Title string
	Items []string
	Date  string
}

func main() {
	// Create a report instance
	report := Report{
		Title: "Daily Report",
		Items: []string{"Item 1", "Item 2", "Item 3"},
		Date:  time.Now().Format("2006-01-02"),
	}

	// Define the HTML template with Sprig functions
	const tpl = `
<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
</head>
<body>
    <h1>{{.Title | title}}</h1>
    <p>Date: {{.Date}}</p>
    <ul>
        {{range .Items}}
        <li>{{. | upper}}</li>
        {{end}}
    </ul>
</body>
</html>
`

	// Parse and execute the template with Sprig functions
	t, err := template.New("report").Funcs(sprig.FuncMap()).Parse(tpl)
	if err != nil {
		panic(err)
	}
	if err := t.Execute(os.Stdout, report); err != nil {
		panic(err)
	}
}
