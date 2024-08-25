package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// PokemonDetails -
func (c *Client) PokemonDetails(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	cacheData, exists := c.cache.Get(url)
	if exists {
		return unMarshalPokemonReturn(cacheData)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, dat)

	return unMarshalPokemonReturn(dat)
}

func unMarshalPokemonReturn(data []byte) (Pokemon, error) {
	locationsResp := Pokemon{}
	err := json.Unmarshal(data, &locationsResp)
	if err != nil {
		return Pokemon{}, err
	}
	return locationsResp, nil
}
