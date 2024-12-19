package main

import "fmt"

// Future Trend: Custom iterator protocol
type MyCustomType []int

func (ct MyCustomType) Iterate() <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range ct {
			c <- v
		}
		close(c)
	}()
	return c
}

func main() {
	myCustomData := MyCustomType{10, 20, 30}
	for value := range myCustomData.Iterate() {
		fmt.Println(value)
	}
}
