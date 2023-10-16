package main

import (
	"fmt"
)

func callbackPokedex(cfg *Config, args ...string) error {
	fmt.Println("Pokemons in pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
