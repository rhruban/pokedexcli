package main

import(
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input			string
		expected	[]string
	}{
		{
			input: "  Hello world  ",
			expected: []string{"hello","world"},
		},
		// TODO add more
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// TODO Check length
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("flail")
			}
		}
	}
}
