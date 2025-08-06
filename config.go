package main

import (
	"github.com/SirVoly/pokedexcli/internal/pokeapi"
	"github.com/SirVoly/pokedexcli/internal/pokecache"
)

type Config struct {
	pokeapiClient        pokeapi.Client
	pokecache            pokecache.Cache
	nextLocationsURL     *string
	previousLocationsURL *string
}
