package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, name string) error {
	if name == "" {
		return errors.New("pokemon name cannot be empty")
	}
	_, exists := cfg.pokedex[name]
	if exists {
		return fmt.Errorf("%s already exists in your pokedex", name)
	}
	pokeDetails, err := cfg.pokeapiClient.PokemonDetails(name)
	if err != nil {
		return err
	}
	fmt.Println("Throwing a Pokeball at " + name + "...")
	chance := rand.Intn(pokeDetails.BaseExperience)
	if chance > 20 {
		return fmt.Errorf("%s escaped!", name)
	}
	fmt.Println(name, " was caught!")
	cfg.pokedex[name] = pokeDetails
	return nil
}
