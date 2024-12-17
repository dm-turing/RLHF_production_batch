package main

import (
	"errors"
	"fmt"
)

func performAction() error {
	return errors.New("action failed")
}

func process() error {
	if err := performAction(); err != nil {
		return fmt.Errorf("process failed: %w", err)
	}
	return nil
}

func main() {
	if err := process(); err != nil {
		fmt.Println(err)
	}
}
