package main

import "fmt"

func commandExplore(cfg *config) error {
	var id string

	if cfg.parameter == "" {
		return fmt.Errorf("Must have location parameter")
	} else {
		id = cfg.parameter
	}
	pokemonResp, err := cfg.pokeapiClient.ExploreLocations(id)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v...\n", cfg.parameter)
	fmt.Println("Found pokemon:")
	for _, pokemon := range pokemonResp.PokemonEncounters {
		fmt.Printf("- %v\n", pokemon.Pokemon.Name)
	}

	return nil
}
