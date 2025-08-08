package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Just1a2Noob/bootdev/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	parameter        string
	pokedex          map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		var parameters string
		if len(words) == 0 {
			continue
		}
		if len(words) >= 2 {
			parameters = words[1]
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			// Checks if parameter exists
			if parameters != "" {
				cfg.parameter = parameters
			}

			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "Finds all pokemons given the location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catches the specified pokemon",
			callback:    commandCatch,
		},
	}
}
