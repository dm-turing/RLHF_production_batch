package main

import (
	"reflect"
	"testing"
)

// This test is to observe if logFunc in the main module doesn't get
// replaced when a local value overrides it.
func Test_doSomething(t *testing.T) {
	var called bool
	logFunc := func(s string) {
		called = true
	}
	reflect.TypeOf(logFunc) // Just a filler to bypass go error 'declared and unused error'
	doSomething()

	if called {
		t.Error("logFunc was not called")
	}
}
