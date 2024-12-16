package main

import (
	"testing"
)

// Function that uses the outdated version of assert for testing
func TestSomething(t *testing.T) {
	var actual = "hello"
	var expected = "hello"

	// Using assert to check equality (this should be present in the latest version)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
