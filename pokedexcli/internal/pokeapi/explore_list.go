package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// Explore Location
func (c *Client) ExploreLocations(id string) (RespIDLocations, error) {

	url := baseURL + "/location-area/" + id

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := RespIDLocations{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return RespIDLocations{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespIDLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespIDLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespIDLocations{}, err
	}

	pokemonResp := RespIDLocations{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespIDLocations{}, err
	}

	c.cache.Add(url, dat)
	return pokemonResp, nil
}
