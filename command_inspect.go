package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("No pokemon name provided")
	} else if len(args) != 1 {
		return errors.New("Please provide only one pokemon name")
	}

	pokemonName := args[0]
	pokemon, exists := cfg.pokedex[pokemonName]

	if !exists {
		return errors.New("You have not caught this pokemon!")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	return nil
}
