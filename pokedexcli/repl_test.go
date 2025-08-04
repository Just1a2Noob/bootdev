package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world     ",
			expected: []string{"hello", "world"},
		},
	}

	results := 0

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of actual != expected")
			results += 1
		}

		for i := range actual {
			word := actual[i]
			expected := c.expected[i]
			if word != expected {
				t.Errorf("Word of actual != expected")
				results += 1
			}
		}
	}
	fmt.Printf("Failed tests: %v", results)
}
