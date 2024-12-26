package main

import (
	"fmt"
	"reflect"
)

func main() {
	callback := getUntrustedCallback()
	if validateCallback(callback) {
		callback()
	} else {
		fmt.Println("Invalid callback")
	}
}

func getUntrustedCallback() func() {
	// Example of untrusted callback source
	return func() {
		fmt.Println("This is a safe callback")
	}
}

func validateCallback(fn func()) bool {
	rv := reflect.ValueOf(fn)
	if rv.Kind() != reflect.Func {
		return false
	}
	numIn := rv.Type().NumIn()
	if numIn != 0 {
		return false
	}
	numOut := rv.Type().NumOut()
	if numOut != 0 {
		return false
	}
	return true
}
