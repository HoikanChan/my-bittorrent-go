package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		{input: "l5:helloi52ed3:foo3:bar5:helloi52eee", expected: []interface{}{"hello", 52, map[string]interface{}{
			"foo":   "bar",
			"hello": 52,
		}}, err: nil},
		{input: "d3:foo3:bar5:helloi52ee", expected: map[string]interface{}{
			"foo":   "bar",
			"hello": 52,
		}, err: nil},

		// Negative test cases
		{input: "invalid", expected: "", err: errInvalidFormat},
		{input: "l5:helloinvalide", expected: "", err: errInvalidFormat},
		{input: "di52e3:bar5:helloi52ee", expected: "", err: errInvalidDictKey},
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
			} else {
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}
