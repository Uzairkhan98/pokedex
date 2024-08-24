module github.com/uzairkhan98/pokedex_repl

go 1.22.4

replace github.com/uzairkhan98/pokeapi v0.0.0 => ./internal/pokeapi

replace github.com/uzairkhan98/pokecache v0.0.0 => ./internal/pokecache

require github.com/uzairkhan98/pokeapi v0.0.0

require github.com/uzairkhan98/pokecache v0.0.0 // indirect
