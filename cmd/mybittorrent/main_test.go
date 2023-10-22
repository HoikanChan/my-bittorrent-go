package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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
		{input: "l5:helloi52ee", expected: []interface{}{"hello", 52}, err: nil},

		// Negative test cases
		{input: "invalid", expected: "", err: errFormat},
		{input: "l5:helloinvalide", expected: "", err: errFormat},
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
			}
			assert.Equal(t, tc.expected, result)
		})
	}
}

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
