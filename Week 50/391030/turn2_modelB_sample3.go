package main

import (
	"github.com/pkg/errors"
)

func outerFunction() error {
	err := innerFunction()
	if err != nil {
		return errors.Wrap(err, "outerFunction failed")
	}
	return nil
}

func innerFunction() error {
	return errors.Wrap(errors.New("innerFunction error"), "some deep function")
}

func main() {
	if err := outerFunction(); err != nil {
		errors.Stack(err)
	}
}
