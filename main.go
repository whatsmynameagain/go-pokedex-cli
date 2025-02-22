package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%v", cleanInput("Hello, World!"))
}

func cleanInput(text string) []string {
	words := strings.Fields(text)

	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return words
}
