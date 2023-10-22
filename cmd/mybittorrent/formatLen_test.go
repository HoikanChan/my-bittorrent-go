package main

import (
	"fmt"
	"testing"
)

func Test_getStringLength(t *testing.T) {
	testCases := []struct {
		input    interface{}
		expected int
		err      error
	}{
		// Positive test cases
		{input: "hello", expected: 7, err: nil},
		{input: 123, expected: 5, err: nil},
		{input: -12345, expected: 8, err: nil},
		{input: []interface{}{"hello", 123}, expected: 14, err: nil},
		{input: map[string]interface{}{
			"foo":   "bar",
			"hello": 52,
		}, expected: 23, err: nil},

		// Negative test cases
		{input: true, expected: -1, err: errGetString},
		{input: nil, expected: -1, err: errGetString},
		{input: 3.14, expected: -1, err: errGetString},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v", tc.input), func(t *testing.T) {
			result, err := getStringLength(tc.input)
			if err != nil {
				if tc.err == nil {
					t.Errorf("Unexpected error: %v", err)
				} else if err.Error() != tc.err.Error() {
					t.Errorf("Expected error %v, but got %v", tc.err, err)
				}
			} else if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}
