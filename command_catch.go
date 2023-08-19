package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("No pokemon name provided")
	} else if len(args) != 1 {
		return errors.New("Please provide only one pokemon name")
	}

	pokemonName := args[0]

	if _, exists := cfg.pokedex[pokemonName]; exists {
		fmt.Printf("%v has already been caught!\n", pokemonName)
		return nil
	}

	data, err := cfg.pokeapiClient.GetPokemon(pokemonName)

	if err != nil {
		log.Fatal(err)
	}

	const threshold = 50
	randNum := rand.Intn(data.BaseExperience)

	fmt.Println(data.BaseExperience, randNum, threshold)

	if randNum > threshold {
		fmt.Printf("Failed to catch %v!\n", pokemonName)
		return nil
	}

	fmt.Printf("Caught %v!\n", pokemonName)
	cfg.pokedex[pokemonName] = data

	return nil
}
