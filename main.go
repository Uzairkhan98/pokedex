package main

import (
	"time"

	"github.com/uzairkhan98/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       pokeapi.Pokedex{},
	}

	startRepl(cfg)
}
