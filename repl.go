package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(" > ")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		command := cleaned[0]
		switch command {
		case "help":
			
		case "exit":
			os.Exit(0)

		default:
			fmt.Println("Invalid command")
		}

	}
}
type cliCommand struct {
	name string
	description string
	callback func()
}
func getCommands () map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Prints the help menu",
			callback: 
		},
		"exit": {
			name: "exit",
			description: "kills the program",
			callback:
		}
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words

}