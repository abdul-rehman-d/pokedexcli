package main

import (
	"errors"
	"fmt"
	"log"
)

func callbackExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("No location areas provided")
	} else if len(args) != 1 {
		return errors.New("Please provide only one location area")
	}

	locationAreaName := args[0]

	data, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Pokemons in %v:\n", data.Name)

	for _, pokemon := range data.PokemonEncounters {
		fmt.Printf(" - %v\n", pokemon.Pokemon.Name)
	}

	return nil
}
