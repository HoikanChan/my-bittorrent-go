package main

import (
	"fmt"
	"testing"
)

func TestDecodeBencode(t *testing.T) {
	// Positive test case: valid bencoded string
	bencodedString := "4:test"
	expected := "test"
	result, err := decodeBencode(bencodedString)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	// Positive test case: valid bencoded number
	bencodedString = "i52e"
	expected = "52"
	result, err = decodeBencode(bencodedString)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	// Negative test case: invalid bencoded string
	bencodedString = "invalid"
	expectedErr := fmt.Errorf("only strings are supported at the moment")
	result, err = decodeBencode(bencodedString)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if err.Error() != expectedErr.Error() {
		t.Errorf("Expected error %v, got %v", expectedErr, err)
	}
}
