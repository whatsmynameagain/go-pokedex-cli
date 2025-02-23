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
			input:    "   hello   w orld   ",
			expected: []string{"hello", "w", "orld"},
		},

		{
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("%d does not match array length %d", len(actual), len(c.expected))
			//this fails the test automatically
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("%s does not match %s", word, expectedWord)
			}
		}

	}
}
