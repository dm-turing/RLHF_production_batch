package main

import (
	"fmt"
	"reflect"
)

func SetField(obj interface{}, fieldName string, value interface{}) error {
	v := reflect.ValueOf(obj).Elem()
	f := v.FieldByName(fieldName)

	if !f.IsValid() {
		return fmt.Errorf("no such field: %s in obj", fieldName)
	}

	if !f.CanSet() {
		return fmt.Errorf("cannot set %s field value", fieldName)
	}

	val := reflect.ValueOf(value)
	if f.Type() != val.Type() {
		return fmt.Errorf("provided value type didn't match obj field type")
	}

	f.Set(val)
	return nil
}

type Person struct {
	Name string
	Age  int
}

func main() {
	p := &Person{}
	if err := SetField(p, "Name", "Bob"); err != nil {
		fmt.Println(err)
	}
	if err := SetField(p, "Age", 30); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Updated Person: %+v\n", p)
}
