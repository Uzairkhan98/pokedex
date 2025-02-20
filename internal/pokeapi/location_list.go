package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	cacheData, exists := c.cache.Get(url)
	if exists {
		return unMarshalShallowReturn(cacheData)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}
	c.cache.Add(url, dat)

	return unMarshalShallowReturn(dat)
}

func unMarshalShallowReturn(data []byte) (RespShallowLocations, error) {
	locationsResp := RespShallowLocations{}
	err := json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}
	return locationsResp, nil
}
