package main

import (
	"fmt"
	"reflect"
)

type Data struct {
	Secret string
}

func main() {
	data := &Data{Secret: "This is a secret"}
	doSomething(data, createCallback(data))
}

func createCallback(data *Data) reflect.Value {
	fn := func() {
		fmt.Println(data.Secret)
	}
	return reflect.ValueOf(fn)
}

func doSomething(data interface{}, callback reflect.Value) {
	// Some operations here...
	callback.Call(nil)
}
