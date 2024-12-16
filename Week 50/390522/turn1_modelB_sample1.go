package main

import (
	"fmt"
	"reflect"
)

func inspectType(value interface{}) {
	v := reflect.ValueOf(value)
	t := reflect.TypeOf(value)

	fmt.Printf("Value: %v\n", v)
	fmt.Printf("Type: %v\n", t)
	fmt.Printf("Kind: %v\n", v.Kind())
	fmt.Printf("Is nil: %+v\n", v.IsNil())
}

func main() {
	var i int = 42
	s := "Hello, World!"
	m := map[string]int{"key": 1337}

	inspectType(i)
	fmt.Println()
	inspectType(s)
	fmt.Println()
	inspectType(m)
}
