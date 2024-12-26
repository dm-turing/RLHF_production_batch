package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"reflect"
)

func executeUntrustedCallback(callback string) {
	fset := token.NewFileSet()
	astFile, err := parser.ParseExpr(fset, callback, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// Create a safe execution environment using reflect.MakeFunc
	fn := reflect.MakeFunc(reflect.TypeOf(func() {}), func([]reflect.Value) []reflect.Value {
		fmt.Println("Hello from untrusted code!")
		return nil
	}).Interface().(func())

	// Convert the ast expression to a reflect.Value and assign it to the variable 'callback'
	callbackValue := reflect.ValueOf(callback)

	// Execute the untrusted code using the reflect.Value
	callbackValue.Call(nil)

}
