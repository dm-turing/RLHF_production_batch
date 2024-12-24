package main

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

type User struct {
	Name    string  `url:"name"`
	Age     int64   `url:"age"`
	IsAdmin bool    `url:"is_admin"`
	Height  float64 `url:"height"`
	Color   string  `url:"color"`
}

// A map to store conversion functions for different types
var convFuncs = map[reflect.Kind]func(string) (interface{}, error){
	reflect.String:  func(s string) (interface{}, error) { return s, nil },
	reflect.Int:     func(s string) (interface{}, error) { return strconv.ParseInt(s, 10, 64) },
	reflect.Int8:    func(s string) (interface{}, error) { return strconv.ParseInt(s, 10, 8) },
	reflect.Int16:   func(s string) (interface{}, error) { return strconv.ParseInt(s, 10, 16) },
	reflect.Int32:   func(s string) (interface{}, error) { return strconv.ParseInt(s, 10, 32) },
	reflect.Int64:   func(s string) (interface{}, error) { return strconv.ParseInt(s, 10, 64) },
	reflect.Bool:    func(s string) (interface{}, error) { return strconv.ParseBool(s) },
	reflect.Float32: func(s string) (interface{}, error) { return strconv.ParseFloat(s, 32) },
	reflect.Float64: func(s string) (interface{}, error) { return strconv.ParseFloat(s, 64) },
}

func parseQuery(query url.Values, result interface{}) error {
	v := reflect.ValueOf(result).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("url")
		if tag == "" {
			continue
		}
		value := query.Get(tag)
		if value == "" {
			continue
		}

		// Look up the conversion function for the field's type
		convFunc, ok := convFuncs[field.Type.Kind()]
		if !ok {
			return fmt.Errorf("unsupported field type: %v", field.Type.Kind())
		}

		// Convert the query value using the conversion function
		convertedValue, err := convFunc(value)
		if err != nil {
			return fmt.Errorf("error converting value for %q: %v", tag, err)
		}

		// Set the converted value on the field
		v.Field(i).Set(reflect.ValueOf(convertedValue))
	}

	return nil
}

func main() {
	query := url.Values{}
	query.Set("name", "Alice")
	query.Set("age", "25")
	query.Set("is_admin", "true")
	query.Set("height", "178.5")
	query.Set("color", "blue")

	var user User
	err := parseQuery(query, &user)
	if err != nil {
		fmt.Printf("Error parsing query: %v\n", err)
		return
	}

	fmt.Printf("User: %+v\n", user)
}
