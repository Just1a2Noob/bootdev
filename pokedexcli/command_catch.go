package main

import (
	"fmt"
	"strings"

	"github.com/Just1a2Noob/bootdev/pokedexcli/internal/pokeapi"
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

	catched := pokeapi.CatchResults(pokemonResp)
	if catched {
		cfg.pokedex[pokemonResp.Name] = pokemonResp
	}

	return nil
}
