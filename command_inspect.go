package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *Config, args ...string) error {

	if len(args) != 1 {
		return errors.New("No pokemon name provided to catch")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("You havent have cought that pokemon!")
	}

	fmt.Printf("Name: %s\n", pokemonName)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)

	fmt.Println("Types:")
	for _, typ := range pokemon.Types {
		fmt.Printf("- %s\n", typ.Type.Name)
	}

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Abilities:")
	for _, ability := range pokemon.Abilities {
		fmt.Printf("- %s\n", ability.Ability.Name)
	}
	return nil
}
