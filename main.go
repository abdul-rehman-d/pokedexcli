package main

import (
	"time"

	"github.com/abdul-rehman-d/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient     pokeapi.Client
	nextLocationAreas *string
	prevLocationAreas *string
	pokedex           map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient:     pokeapi.NewClient(time.Minute * 5),
		nextLocationAreas: nil,
		prevLocationAreas: nil,
		pokedex:           make(map[string]pokeapi.Pokemon),
	}

	startRepl(&cfg)
}
