package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *Config, args ...string) error {

	if len(args) != 1 {
		return errors.New("No pokemon name provided to catch")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	threshold := 50 // it should be your level * 10 or something like that but we don't have levels that is why it is 50 by default
	randNum := rand.Intn(pokemon.BaseExperience)
	// fmt.Println(pokemon.BaseExperience, randNum, threshold)
	fmt.Printf("throwing a pokeball to %s", pokemonName)
	if randNum > threshold {
		return fmt.Errorf("Failed to catch %s"+"!", pokemonName)
	}

	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("%s has cought!\n", pokemonName)

	return nil
}
