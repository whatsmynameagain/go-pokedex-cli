package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	inScanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		for inScanner.Scan() {
			userInput := inScanner.Text()
			userInput = strings.ToLower(userInput)
			words := strings.Fields(userInput)
			fmt.Printf("Your command was: %s\n", words[0])
			fmt.Print("Pokedex > ")
		}
		if err := inScanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
		}
	}

}

func cleanInput(text string) []string {
	words := strings.Fields(text)

	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return words
}
