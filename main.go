package main

import (
	"time"

	"github.com/uzairkhan98/pokeapi"
	"github.com/uzairkhan98/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		cache:         pokecache.NewCache(5 * time.Minute)}

	startRepl(cfg)
}
