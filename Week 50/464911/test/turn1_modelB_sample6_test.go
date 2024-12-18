package main

import (
	"fmt"
	"testing"
)

func TestSomeFunction(t *testing.T) {
	result := SomeFunction()
	expectedResult := fmt.Errorf("%v", "SomeFunction failed: Some error happened")
	if result.Error() != expectedResult.Error() {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}
