package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	availableCommands := getCommands()

	for {
		fmt.Printf("> ")
		scanner.Scan()

		str := scanner.Text()
		cleaned := cleanInput(str)

		if len(cleaned) == 0 {
			continue
		}

		commandFromUser := cleaned[0]

		command, exists := availableCommands[commandFromUser]

		if !exists {
			fmt.Println("Invalid command")
			continue
		}

		command.callback()
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func()
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays this help menu",
			callback:    printHelpMenu,
		},
		"exit": {
			name:        "exit",
			description: "Exits program",
			callback: func() {
				os.Exit(0)
			},
		},
	}
}

func printHelpMenu() {
	fmt.Println("Welcome to Pokedex CLI")
	fmt.Println("Available Commands:")

	for _, cmd := range getCommands() {
		fmt.Printf("\t%v\t\t%v\n", cmd.name, cmd.description)
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)

	return words
}
