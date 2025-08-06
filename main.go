package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		fmt.Println("Your command was:", cleanInput(input)[0])
	}
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
