package main

import (
	"errors"
	"fmt"
	"math/rand"
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
		"catch": {
			name:        "catch <pokemon>",
			description: "You try and catch the Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "You check one of your caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "You check all of your caught Pokemon",
			callback:    commandPokedex,
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

func commandCatch(cfg *Config) error {
	if len(cfg.commandArgs) != 1 {
		return fmt.Errorf("catch takes exactly one argument: the name of the pokemon")
	}

	pokemonName := cfg.commandArgs[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	success := succeedCatch(pokemon.BaseExperience)

	if success {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.caughtPokemon[pokemonName] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}

func succeedCatch(difficulty int) bool {
	if difficulty < 0 {
		difficulty = 0
	}
	if difficulty > 400 {
		difficulty = 400
	}
	prob := 1.0 - float64(difficulty)/400.0 // 0 = 100%, 400 = 0%
	return rand.Float64() < prob
}

func commandInspect(cfg *Config) error {
	if len(cfg.commandArgs) != 1 {
		return fmt.Errorf("catch takes exactly one argument: the name of the pokemon")
	}

	pokemonName := cfg.commandArgs[0]

	pokemon, caught := cfg.caughtPokemon[pokemonName]

	if !caught {
		return fmt.Errorf("you have not caught %s", pokemonName)
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("\t- %s\n", t.Type.Name)
	}

	return nil
}

func commandPokedex(cfg *Config) error {
	fmt.Println("Your Pokedex:")

	if len(cfg.caughtPokemon) == 0 {
		return fmt.Errorf("you have not caught any pokemon yet")
	}

	for pokemonName := range cfg.caughtPokemon {
		fmt.Println("\t- ", pokemonName)
	}

	return nil
}
