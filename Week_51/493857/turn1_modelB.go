package main

import (
	"fmt"
)

type MyTime struct {
	Hour   int
	Minute int
	Second int
}

func (t MyTime) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%02d:%02d:%02d", t.Hour, t.Minute, t.Second)
		} else {
			fmt.Fprintf(s, "%02d:%02d", t.Hour, t.Minute)
		}
	case 's':
		fmt.Fprintf(s, "%02d:%02d:%02d", t.Hour, t.Minute, t.Second)
	default:
		fmt.Fprintf(s, "%%!%c(MyTime=%+v)", verb, t)
	}
}
func main() {
	t := MyTime{Hour: 10, Minute: 30, Second: 15}
	fmt.Println(t)         // Output: 10:30
	fmt.Printf("%+v\n", t) // Output: 10:30:15
	fmt.Printf("%s\n", t)  // Output: 10:30:15
	fmt.Printf("%q\n", t)  // Output: "10:30:15"
}
