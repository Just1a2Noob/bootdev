package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World")
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
