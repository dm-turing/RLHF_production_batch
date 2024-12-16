package main

import (
	"fmt"
	"reflect"
)

func printValue(v interface{}) {
	rv := reflect.ValueOf(v)
	fmt.Println(rv.Interface())
}

func main() {
	printValue(42)      // Output: 42
	printValue("hello") // Output: hello
}
