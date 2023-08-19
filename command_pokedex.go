package main

import "fmt"

func callbackPokedex(cfg *config, _ ...string) error {

	if len(cfg.pokedex) == 0 {
		fmt.Println("You haven't caught any pokemons yet!")
		return nil
	}

	fmt.Println("Caught Pokemons:")
	for name := range cfg.pokedex {
		fmt.Printf("- %s\n", name)
	}

	return nil
}
