package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commands map[string]cliCommand

func main() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())

		command, exists := commands[input[0]]
		if !exists {
			fmt.Println("Unknown command")
		} else {
			err := command.callback()
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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for _, c := range commands {
		fmt.Printf("\t%s: %s\n", c.name, c.description)
	}
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}
