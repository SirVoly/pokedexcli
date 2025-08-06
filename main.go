package main

import (
	"time"

	"github.com/SirVoly/pokedexcli/internal/pokeapi"
	"github.com/SirVoly/pokedexcli/internal/pokecache"
)

func main() {

	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokecache.NewCache(5 * time.Second)

	cfg := &Config{
		pokeapiClient: pokeClient,
		pokecache:     pokeCache,
	}

	LoadCommands()

	LaunchREPL(cfg)
}
