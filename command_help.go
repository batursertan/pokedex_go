package main

import "fmt"

func callbackHelp(cfg *Config) error {
	fmt.Println("Welcome to the pokedex help menu!")
	fmt.Println("Here are your avaliable commands:")

	avaliablecommands := getCommands()
	for _, cmd := range avaliablecommands {
		fmt.Printf(" - %s: %s \n", cmd.name, cmd.description)

	}
	fmt.Println("")
	return nil
}
