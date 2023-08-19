package main

import (
	"fmt"
	"log"
)

func callbackMap(cfg *config) error {
	data, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreas)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location areas:")

	for _, location := range data.Results {
		fmt.Printf(" - %v\n", location.Name)
	}

	cfg.nextLocationAreas = data.Next
	cfg.prevLocationAreas = data.Previous

	return nil
}
