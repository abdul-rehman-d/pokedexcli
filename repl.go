package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
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

		if command, exists := availableCommands[cleaned[0]]; exists {
			command.callback(cfg)
		} else {
			fmt.Println("Invalid command")
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays this help menu",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits program",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "List of locations",
			callback:    callbackMap,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)

	return words
}
