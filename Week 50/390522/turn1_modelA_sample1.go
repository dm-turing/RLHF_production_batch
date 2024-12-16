package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func describe(v interface{}) {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	fmt.Printf("Type: %s\n", typ)
	fmt.Printf("Value: %v\n", val)

	if val.Kind() == reflect.Struct {
		fmt.Println("Struct Fields:")
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fieldType := typ.Field(i)
			fmt.Printf("Field: %s, Type: %s, Value: %v\n", fieldType.Name, fieldType.Type, field.Interface())
		}
	}
}

func main() {
	p := Person{Name: "Alice", Age: 25}
	describe(p)
}
