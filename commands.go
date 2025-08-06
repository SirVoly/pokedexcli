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
		"explore": {
			name:        "explore <location-area>",
			description: "Displays the names of the pokemon on a specific location in the Pokemon world",
			callback:    commandExplore,
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
	pokeAPIPageResponse, err := cfg.pokeapiClient.GetLocationsList(cfg.nextLocationsURL)

	if err != nil {
		return err
	}

	cfg.nextLocationsURL = pokeAPIPageResponse.Next
	cfg.previousLocationsURL = pokeAPIPageResponse.Previous

	for _, location := range pokeAPIPageResponse.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *Config) error {
	// displays the names of the previous 20 location areas in the Pokemon world
	if cfg.previousLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	pokeAPIPageResponse, err := cfg.pokeapiClient.GetLocationsList(cfg.previousLocationsURL)

	if err != nil {
		return err
	}

	cfg.nextLocationsURL = pokeAPIPageResponse.Next
	cfg.previousLocationsURL = pokeAPIPageResponse.Previous

	for _, location := range pokeAPIPageResponse.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandExplore(cfg *Config) error {

	if len(cfg.commandArgs) != 1 {
		return fmt.Errorf("explore takes exactly one argument: the name of the location area")
	}

	exploredLocation := cfg.commandArgs[0]

	location, err := cfg.pokeapiClient.GetPokemonsAtLocationList(exploredLocation)

	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", exploredLocation)

	if len(location.PokemonEncounters) > 0 {
		fmt.Println("Found Pokemon:")
		for _, pokemon := range location.PokemonEncounters {
			fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
		}
	} else {
		fmt.Println("No Pokemon were found.")
	}

	return nil
}
