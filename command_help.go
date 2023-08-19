package main

import "fmt"

func callbackHelp(_ *config, _ ...string) error {
	fmt.Println("Welcome to Pokedex CLI")
	fmt.Println("Available Commands:")

	for name, cmd := range getCommands() {
		fmt.Printf("\t%v\t\t%v\n", name, cmd.description)
	}

	return nil
}
