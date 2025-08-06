package main

import (
	"errors"
	"fmt"
	"os"
)

var commands map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *Config) error
}

func LoadCommands() {
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
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
	}
}

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, c := range commands {
		fmt.Printf("\t%s: %s\n", c.name, c.description)
	}
	return nil
}

func commandMap(cfg *Config) error {
	// displays the names of the next 20 location areas in the Pokemon world
	response, err := cfg.pokeapiClient.GetLocationsList(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = response.Next
	cfg.previousLocationsURL = response.Previous

	for _, location := range response.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *Config) error {
	// displays the names of the previous 20 location areas in the Pokemon world
	if cfg.previousLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	response, err := cfg.pokeapiClient.GetLocationsList(cfg.previousLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = response.Next
	cfg.previousLocationsURL = response.Previous

	for _, location := range response.Results {
		fmt.Println(location.Name)
	}
	return nil
}
