package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	bencodedValue := "d5:hello5:worlde"
	os.Args = []string{"main", "decode", bencodedValue}
	expectedOutput := map[string]interface{}{
		"hello": "world",
	}
	decoded, _,err := decodeBencode(bencodedValue)
	if err != nil {
		t.Errorf("decodeBencode failed with error: %s", err.Error())
	}
	assert.Equal(t, decoded, expectedOutput)
}

func TestInfo(t *testing.T) {
	// Positive test case for "info" command
	torrentContent := "d5:hello5:worlde"
	tmpFile, err := ioutil.TempFile("", "torrent")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	if _, err := tmpFile.Write([]byte(torrentContent)); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	if err := tmpFile.Close(); err != nil {
		t.Fatalf("Failed to close temporary file: %v", err)
	}
	os.Args = []string{"main", "info", "../../sample.torrent"}
	main()

	// Negative test case for unknown command
	// os.Args = []string{"main", "unknown"}
	// main()
}
