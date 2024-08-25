package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListPokemons -
func (c *Client) ListPokemons(pageURL string) (RespDeepLocations, error) {
	url := baseURL + "/location-area/" + pageURL

	cacheData, exists := c.cache.Get(url)
	if exists {
		return unMarshalDeepReturn(cacheData)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespDeepLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDeepLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDeepLocations{}, err
	}
	c.cache.Add(url, dat)

	return unMarshalDeepReturn(dat)
}

func unMarshalDeepReturn(data []byte) (RespDeepLocations, error) {
	locationsResp := RespDeepLocations{}
	err := json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespDeepLocations{}, err
	}
	return locationsResp, nil
}
