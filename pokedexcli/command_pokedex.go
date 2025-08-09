package main

import "fmt"

func commandPokedex(cfg *config) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("You haven't captured any Pokemon's")
		return nil
	}

	fmt.Println("Your Pokedex:")

	for key := range cfg.pokedex {
		fmt.Printf("- %s\n", key)
	}

	return nil
}
