package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func start() {

	inScanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		inScanner.Scan()
		words := cleanInput(inScanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		if command, exists := getCommands()[commandName]; exists {
			err := command.callback()
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
	words := strings.Fields(text)

	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}
