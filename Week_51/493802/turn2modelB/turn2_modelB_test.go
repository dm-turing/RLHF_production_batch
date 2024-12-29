package main

import (
	"testing"
)

func Test_someFunction(t *testing.T) {
	testCases := []struct {
		a        int
		b        int
		expected int
	}{
		{1, 2, 3},
		{3, 4, 7},
		{-1, -2, -3},
	}

	for _, testCase := range testCases {
		result := someFunction(testCase.a, testCase.b)
		if result != testCase.expected {
			t.Errorf("someFunction(%d, %d) returned %d, expected %d", testCase.a, testCase.b, result, testCase.expected)
		}
	}
}
