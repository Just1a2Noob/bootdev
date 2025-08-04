package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commands map[string]cliCommand

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		for scanner.Scan() {
			response := scanner.Text()
			text := cleanInput(response)

			for _, word := range text {
				if _, ok := commands[word]; !ok {
					continue
				}
				commands[word].callback()
			}
			break
		}
	}
}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)

	words := strings.Fields(trimmed)

	results := make([]string, len(words))
	for i, word := range words {
		results[i] = strings.ToLower(word)
	}

	return results
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// CLI commands
func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Shows commands name and functionality",
			callback:    commandHelp,
		},
	}
}

// CLI Functions
func commandExit() error {
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for _, command := range commands {
		fmt.Printf("%v : %v\n", command.name, command.description)
	}
	return nil
}
