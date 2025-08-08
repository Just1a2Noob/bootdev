package main

import "fmt"

func commandExplore(cfg *config) error {
	var id string

	if cfg.parameter == "" {
		return fmt.Errorf("Must have location parameter")
	} else {
		id = cfg.parameter
	}
	pokemonResp, err := cfg.pokeapiClient.ExploreLocations(&id)
	if err != nil {
		return err
	}

	for _, pokemon := range pokemonResp.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
