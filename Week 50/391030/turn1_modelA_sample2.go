package main

import (
	"errors"
	"fmt"
)

func doSomething() error {
	return errors.New("original error")
}

func main() {
	err := doSomething()
	if err != nil {
		err = fmt.Errorf("doSomething failed: %w", err)
		fmt.Println(err)

		if errors.Is(err, errors.New("original error")) {
			fmt.Println("The original error occurred")
		}
	}
}
