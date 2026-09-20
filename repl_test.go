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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  whywhy why ",
			expected: []string{"whywhy", "why"},
		},
		{
			input:    "  hel lo  wor ld  ",
			expected: []string{"hel", "lo", "wor", "ld"},
		},
	}

	for _, c := range cases {
		actual := clean_input(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of expected output and actual output don't match")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Word output does not match expected word")
			}
		}

	}

}
