package main

import (
	"time"

	"github.com/SirVoly/pokedexcli/internal/pokeapi"
)

func main() {

	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &Config{
		pokeapiClient: pokeClient,
	}

	LoadCommands()

	LaunchREPL(cfg)
}
