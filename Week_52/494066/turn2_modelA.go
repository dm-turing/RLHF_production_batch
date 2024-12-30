package main

import (
	"fmt"
	"html/template"
	"os"
	"time"
)

func fullName(first string, last string) string {
	return first + " " + last
}

func formatPrice(amount float64) string {
	return fmt.Sprintf("$%.2f", amount)
}

func main() {
	name := "Dana"
	message := fmt.Sprintf("Hello, %v! Today is %s, %d%s, and your percentage is %.2f%%.", name, time.Now().Format("2006-01-02"), time.Now().Year(), ".", 85.5)
	fmt.Println(message)
	tmpl := template.Must(template.New("").Parse("Hello, {{ .Name }}! The score is {{ .Score }}/{{ .MaxScore }}."))

	data := struct {
		Name     string
		Score    int
		MaxScore int
	}{
		Name:     "Eric",
		Score:    95,
		MaxScore: 100,
	}

	if err := tmpl.Execute(os.Stdout, data); err != nil {
		panic(err)
	}

	tmpl, _ = template.Must(template.New("").Funcs(template.FuncMap{
		"FullName":    fullName,
		"FormatPrice": formatPrice,
	}), nil).Parse("{{ FullName .First .Last }} is a buyer of {{ FormatPrice .ProductPrice }}.")

	data1 := struct {
		First        string
		Last         string
		ProductPrice float64
	}{
		First:        "Frank",
		Last:         "Grant",
		ProductPrice: 19.99,
	}

	if err := tmpl.Execute(os.Stdout, data1); err != nil {
		panic(err)
	}
}
