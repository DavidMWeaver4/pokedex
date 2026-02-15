package main

import (
	"testing"
)
func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		// Basic whitespace trimming and lowercase
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello\nworld",
			expected: []string{"hello", "world"},
		},

		// Multiple spaces and tabs
		{
			input:    "Go\tis\tawesome",
			expected: []string{"go", "is", "awesome"},
		},
		{
			input:    "   spaced    out   words   ",
			expected: []string{"spaced", "out", "words"},
		},

		// Empty string
		{
			input:    "",
			expected: []string{},
		},

		// Only whitespace
		{
			input:    "    \t\n   ",
			expected: []string{},
		},

		// Mixed case
		{
			input:    "ThIs Is A TeSt",
			expected: []string{"this", "is", "a", "test"},
		},

		// Single word
		{
			input:    "Word",
			expected: []string{"word"},
		},

		// Newlines and multiple whitespace types
		{
			input:    "line1\nline2\tline3  line4",
			expected: []string{"line1", "line2", "line3", "line4"},
		},
	}

	for _, c := range cases {
			actual := cleanInput(c.input)

			// First, check if lengths match
			if len(actual) != len(c.expected) {
				t.Errorf("cleanInput(%q) returned %d words; expected %d", c.input, len(actual), len(c.expected))
				continue // skip element comparison if lengths differ
			}

			// Compare each element
			for i := range actual {
				if actual[i] != c.expected[i] {
					t.Errorf("cleanInput(%q)[%d] = %q; expected %q", c.input, i, actual[i], c.expected[i])
				}
			}
		}
}
