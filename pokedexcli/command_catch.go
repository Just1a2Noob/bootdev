package main

import (
	"fmt"
	"strings"
)

func commandCatch(cfg *config) error {
	if cfg.parameter == "" {
		return fmt.Errorf("Must specify pokemon name")
	}

	text := strings.ToLower(cfg.parameter)
	pokemonResp, err := cfg.pokeapiClient.PokemonStats(text)
	if err != nil {
		return err
	}

	catched := cfg.pokedex.CatchResults(pokemonResp)
	if catched {
		cfg.pokedex.Add(&pokemonResp)
	}

	return nil
}
