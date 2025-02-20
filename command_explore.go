package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, str string) error {
	if str == "" {
		return errors.New("please provide a location name or ID")
	}
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
