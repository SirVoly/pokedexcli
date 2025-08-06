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
			input:    "  HELLO  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  HeLLo  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "    ",
			expected: []string{},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "hello",
			expected: []string{"hello"},
		},
		{
			input:    "   hello   ",
			expected: []string{"hello"},
		},
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello\tworld",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello\nworld",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello  \n  \t  world",
			expected: []string{"hello", "world"},
		},
		{
			input:    " hello  world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello  world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello12 ",
			expected: []string{"hello12"},
		},
		{
			input:    "hello-2  ",
			expected: []string{"hello-2"},
		},
		{
			input:    "hello?  ",
			expected: []string{"hello?"},
		},
		{
			input:    "CAFÉ résumé",
			expected: []string{"café", "résumé"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		// Check length match
		if len(actual) != len(c.expected) {
			t.Errorf(`
Test Failed:
	input: %v
	expected: %v
	actual: %v
`, c.input, c.expected, actual)
			continue
		}

		// Check each word in the slice
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf(`
Test Failed:
	input: %v
	expected: %v
	actual: %v
`, c.input, c.expected, actual)
				break
			}
		}
	}
}
