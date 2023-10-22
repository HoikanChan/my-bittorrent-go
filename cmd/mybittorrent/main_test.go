package main

import (
	"fmt"
	"testing"
)

func TestDecodeBencode(t *testing.T) {
	testCases := []struct {
		input    string
		expected interface{}
		err      error
	}{
		// Positive test cases
		{input: "5:hello", expected: "hello", err: nil},
		{input: "i123e", expected: 123, err: nil},
		{input: "i-52e", expected: -52, err: nil},
		{input: "4:test", expected: "test", err: nil},

		// Negative test cases
		{input: "invalid", expected: "", err: fmt.Errorf("only strings are supported at the moment")},
		{input: "l123e", expected: "", err: fmt.Errorf("only strings are supported at the moment")},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result, err := decodeBencode(tc.input)
			if err != nil {
				if tc.err == nil {
					t.Errorf("Unexpected error: %v", err)
				} else if err.Error() != tc.err.Error() {
					t.Errorf("Expected error %v, but got %v", tc.err, err)
				}
			} else if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}