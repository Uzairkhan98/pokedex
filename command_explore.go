package main

import (
	"fmt"
)

func commandExplore(cfg *config, str string) error {
	fmt.Println("Exploring " + str + "...")
	exploreResp, err := cfg.pokeapiClient.ListPokemons(str)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range exploreResp.PokemonEncounters {
		fmt.Println(" - " + pokemon.Pokemon.Name)
	}
	return nil
}
