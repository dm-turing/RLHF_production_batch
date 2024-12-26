package main

import (
	"fmt"
	"sort"
	"strings"
)

type Histogram struct {
	Data map[string]int
}

func (h Histogram) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%s", h.String())
		} else {
			fmt.Fprintf(s, "%s", h.SimpleString())
		}
	case 's':
		fmt.Fprintf(s, "%s", h.SimpleString())
	default:
		fmt.Fprintf(s, "%%!%c(Histogram=%+v)", verb, h)
	}
}
func (h Histogram) SimpleString() string {
	var buf strings.Builder
	for key, value := range h.Data {
		buf.WriteString(fmt.Sprintf("%s: %d\n", key, value))
	}
	return buf.String()
}
func (h Histogram) String() string {
	var buf strings.Builder
	maxValue := 0
	for _, value := range h.Data {
		if value > maxValue {
			maxValue = value
		}
	}
	for _, key := range sortedKeys(h.Data) {
		value := h.Data[key]
		bar := strings.Repeat("#", value*50/maxValue)
		buf.WriteString(fmt.Sprintf("%s: %d |%s|\n", key, value, bar))
	}
	return buf.String()
}
func sortedKeys(data map[string]int) []string {
	var keys []string
	for key := range data {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
func main() {
	data := map[string]int{
		"Red":    25,
		"Green":  10,
		"Blue":   5,
		"Yellow": 20,
	}
	h := Histogram{Data: data}
	fmt.Println(h) // Output: Red: 25 |##############################|
	// Green: 10 |##########|
	// Blue:  5 |####|
	// Yellow: 20 |####################|
	fmt.Printf("%+v\n", h) // Output: Red: 25 |##############################|
	// Green: 10 |##########|
	// Blue:  5 |####|
	// Yellow: 20 |####################|
	fmt.Printf("%s\n", h) // Output: Red: 25
	// Green: 10
	// Blue:  5
	// Yellow: 20
}
