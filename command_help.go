package main

import "fmt"

func callbackHelp() {
	fmt.Println("Welcome to Pokedex CLI")
	fmt.Println("Available Commands:")

	for _, cmd := range getCommands() {
		fmt.Printf("\t%v\t\t%v\n", cmd.name, cmd.description)
	}
}
