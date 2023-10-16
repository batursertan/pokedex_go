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

	threshold := pokemon.BaseExperience * 1 / 2
	randNum := rand.Intn(pokemon.BaseExperience)
	// fmt.Println(pokemon.BaseExperience, randNum, threshold)
	fmt.Printf("throwing a pokeball to %s", pokemonName)
	if randNum < threshold {
		return fmt.Errorf("Failed to catch %s"+"!", pokemonName)
	}

	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("%s has cought!\n", pokemonName)

	return nil
}
