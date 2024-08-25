package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, str string) error {
	if str == "" {
		return errors.New("pokemon name cannot be empty")
	}
	_, exists := cfg.pokedex[str]
	if exists {
		return fmt.Errorf("%s already exists in your pokedex", str)
	}
	fmt.Println("Throwing a Pokeball at " + str + "...")
	pokeDetails, err := cfg.pokeapiClient.PokemonDetails(str)
	if err != nil {
		return err
	}
	chance := rand.Intn(pokeDetails.BaseExperience)
	if chance > 20 {
		return fmt.Errorf("%s escaped!", str)
	}
	fmt.Println(str, " was caught!")
	cfg.pokedex[str] = pokeDetails
	return nil
}
