package main

import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	name := "Alice"
	greeting := "Hello, " + name[:] + "!"
	fmt.Println(greeting)
	input := "Hello World"
	formatted := input[0:5] + "world"
	fmt.Println(formatted)
	name = "Bob"
	age := 30
	formatted = fmt.Sprintf("Hello, my name is %s and I am %d years old.", name, age)
	fmt.Println(formatted)
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		panic(err)
	}
	person := Person{Name: "Charlie", Age: 25}
	if err := tmpl.Execute(os.Stdout, person); err != nil {
		panic(err)
	}
}
