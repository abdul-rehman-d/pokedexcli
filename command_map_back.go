package main

import (
	"fmt"
	"log"
)

func callbackMapBack(cfg *config, _ ...string) error {
	data, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreas)

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
