package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func modifyFields(p *Person) {
	v := reflect.ValueOf(p).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fv := v.Field(i)

		fmt.Printf("Field %d: %v (%v)\n", i, field.Name, fv)

		// Modify the field if it's a string or int
		switch fv.Kind() {
		case reflect.String:
			fv.SetString(field.Name + " Modified")
		case reflect.Int:
			fv.SetInt(fv.Int() + 1)
		}
	}
}

func main() {
	person := &Person{Name: "Alice", Age: 30}
	fmt.Println("Before modification:")
	fmt.Printf("%+v\n", person)

	modifyFields(person)

	fmt.Println("\nAfter modification:")
	fmt.Printf("%+v\n", person)
}
