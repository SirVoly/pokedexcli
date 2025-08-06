package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

// Split user input based on whitespaces.
// Lowercases and trims all words.
func cleanInput(text string) []string {
	var result []string

	for word := range strings.FieldsSeq(text) {
		trimmedWord := strings.TrimSpace(word)
		if trimmedWord != "" {
			lowercaseWord := strings.ToLower((trimmedWord))
			result = append(result, lowercaseWord)
		}
	}

	return result
}
