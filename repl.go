package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("> ")
		scanner.Scan()

		str := scanner.Text()
		fmt.Printf("You said: %v\n", str)
	}
}
