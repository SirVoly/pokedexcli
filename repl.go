package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LaunchREPL(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())

		command, exists := commands[input[0]]
		if !exists {
			fmt.Println("Unknown command")
		} else {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		}

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
