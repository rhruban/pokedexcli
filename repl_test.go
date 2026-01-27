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
		{
			input: "Bulbasaur Charmander    PICKACHU",
			expected: []string{"bulbasaur","charmander","pickachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("not same length")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("not same words")
			}
		}
	}
}
