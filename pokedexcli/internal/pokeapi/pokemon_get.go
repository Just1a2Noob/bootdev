package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
)

func (c *Client) PokemonStats(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonResp := Pokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return pokemonResp, err
	}

	c.cache.Add(url, dat)
	return pokemonResp, err
}

func CatchResults(pokemon Pokemon) bool {
	name := strings.ToLower(pokemon.Name)
	success_rate := float64(pokemon.BaseExperience) * 0.45
	results := float64(rand.Intn(pokemon.BaseExperience))

	if success_rate >= results {
		fmt.Printf("%s was caught!\n", name)
		return true
	}
	fmt.Printf("%s escaped!\n", name)

	return false
}
