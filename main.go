package main

import "github.com/abdul-rehman-d/pokedexcli/internal/pokeapi"

type config struct {
	pokeapiClient     pokeapi.Client
	nextLocationAreas *string
	prevLocationAreas *string
}

func main() {
	cfg := config{
		pokeapiClient:     pokeapi.NewClient(),
		nextLocationAreas: nil,
		prevLocationAreas: nil,
	}

	startRepl(&cfg)
}
