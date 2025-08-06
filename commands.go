package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/SirVoly/pokedexcli/internal/pokeapi"
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
	url := pokeapi.BaseURL + "/location-area"
	if cfg.nextLocationsURL != nil {
		url = *cfg.nextLocationsURL
	}

	var dat []byte
	var pokeAPIPageResponse pokeapi.PokeAPILocationResponse

	dat, cached := cfg.pokecache.Get(url)
	if !cached {
		fmt.Println("No Cache, getting fresh data!")
		new_dat, err := cfg.pokeapiClient.GetLocationsList(url)
		if err != nil {
			return err
		}
		cfg.pokecache.Add(url, new_dat)
		dat = new_dat
	}

	if err := json.Unmarshal(dat, &pokeAPIPageResponse); err != nil {
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
	url := *cfg.previousLocationsURL

	var dat []byte
	var pokeAPIPageResponse pokeapi.PokeAPILocationResponse

	dat, cached := cfg.pokecache.Get(url)
	if !cached {
		fmt.Println("No Cache, getting fresh data!")
		new_dat, err := cfg.pokeapiClient.GetLocationsList(url)
		if err != nil {
			return err
		}
		cfg.pokecache.Add(url, new_dat)
		dat = new_dat
	}

	if err := json.Unmarshal(dat, &pokeAPIPageResponse); err != nil {
		return err
	}

	cfg.nextLocationsURL = pokeAPIPageResponse.Next
	cfg.previousLocationsURL = pokeAPIPageResponse.Previous

	for _, location := range pokeAPIPageResponse.Results {
		fmt.Println(location.Name)
	}
	return nil
}
