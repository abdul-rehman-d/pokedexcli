package main

import (
	"fmt"

	"github.com/abdul-rehman-d/pokedexcli/internal/pokeapi"
)

func callbackMap() error {
	client := pokeapi.NewClient()
	data, err := client.ListLocationAreas()

	if err != nil {
		return err
	}

	fmt.Println("Location areas:")

	for _, location := range data.Results {
		fmt.Printf(" - %v\n", location.Name)
	}

	return nil
}
