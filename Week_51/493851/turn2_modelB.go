package main

import "fmt"

type TemplateData struct {
	Title   string
	Message string
}

func RenderTemplate(data TemplateData) string {
	html := `
        <html>
            <head>
                <title>%s</title>
            </head>
            <body>
                <h1>%s</h1>
            </body>
        </html>
    `
	return fmt.Sprintf(html, data.Title, data.Message)
}

func main() {
	data := TemplateData{
		Title:   "Welcome",
		Message: "Hello, world!",
	}
	fmt.Println(RenderTemplate(data))
}
