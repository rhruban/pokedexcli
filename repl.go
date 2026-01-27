package main

import(
	"strings"
)

func cleanInput(text string) []string {
	var words []string
	text = strings.ToLower(text)
	words = strings.Fields(text)
	return words
}
