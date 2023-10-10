package main

import "github.com/batursertan/pokedex_go/internal/pokeapi"

type Config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {

	cfg := Config{
		pokeapiClient: pokeapi.NewClient(),
	}

	startRepl(&cfg)
}
