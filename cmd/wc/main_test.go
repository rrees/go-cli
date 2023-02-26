package main

import (
	"bytes"
	"testing"
)

func TestCountingWords(t *testing.T) {

	tests := []struct {
		Input         string
		ExpectedCount int
	}{
		{"word1 word2 word3 word4", 4},
		{"Word1 word2. Word3 word4.", 4},
		{"Word", 1},
	}

	for _, testData := range tests {

		result := count(bytes.NewBufferString(testData.Input))

		if result != testData.ExpectedCount {
			t.Errorf("Expected %d words but counted %d", testData.ExpectedCount, result)
		}
	}
}
