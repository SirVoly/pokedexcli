package main

import "github.com/SirVoly/pokedexcli/internal/pokeapi"

type Config struct {
	pokeapiClient        pokeapi.Client
	nextLocationsURL     *string
	previousLocationsURL *string
}
