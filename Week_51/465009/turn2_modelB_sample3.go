// Go code using fmt.Fprintf
package main

import (
	"bytes"
	"fmt"
)

func main() {
	data := []int{1, 3, 5, 7, 9}
	var result bytes.Buffer
	fmt.Fprintf(&result, "Elements: [")
	for i, value := range data {
		fmt.Fprintf(&result, "%d", value)
		if i < len(data)-1 {
			fmt.Fprintf(&result, ", ")
		}
	}
	fmt.Fprintf(&result, "]")
	fmt.Println(result.String())
}
