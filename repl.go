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
			err := command.callback(cfg, cleaned[1:]...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Invalid command")
		}
	}
}

type cliCommand struct {
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			description: "Displays this help menu",
			callback:    callbackHelp,
		},
		"map": {
			description: "List of locations",
			callback:    callbackMap,
		},
		"mapb": {
			description: "List of prev locations",
			callback:    callbackMapBack,
		},
		"explore": {
			description: "Explore location",
			callback:    callbackExplore,
		},
		"clear": {
			description: "Clear screen",
			callback:    callbackClear,
		},
		"exit": {
			description: "Exits program",
			callback:    callbackExit,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)

	return words
}
