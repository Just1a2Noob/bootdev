package main

import (
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
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    " HeLLo wOrlD ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of actual != expected")
		}

		for i := range actual {
			word := actual[i]
			expected := c.expected[i]
			if word != expected {
				t.Errorf("Word of actual != expected")
			}
		}
	}
}
