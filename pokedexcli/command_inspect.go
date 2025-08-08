package main

import (
	"fmt"
	"strings"
)

func commandInspect(cfg *config) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("The Pokedex is empty")
	}

	_, ok := cfg.pokedex[cfg.parameter]
	if ok != true {
		fmt.Println("you have not caught that pokemon")
	}

	text := strings.ToLower(cfg.parameter)

	pokemonResp, err := cfg.pokeapiClient.PokemonStats(text)
	if err != nil {
		return err
	}

	fmt.Printf("Name: %s\n", text)
	fmt.Printf("Height: %v\n", pokemonResp.Height)
	fmt.Printf("Weight: %v\n", pokemonResp.Weight)

	fmt.Println("Stats:")

	for _, val := range pokemonResp.Stats {
		fmt.Printf(" -%s: %v\n", val.Stat.Name, val.BaseStat)
	}

	fmt.Println("Types:")

	for _, val := range pokemonResp.Types {
		fmt.Println(val.Type.Name)
	}
	return nil
}
