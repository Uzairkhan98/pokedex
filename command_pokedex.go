package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, str string) error {
	if len(cfg.pokedex) == 0 {
		return errors.New("No pokemon in the pokedex")
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
