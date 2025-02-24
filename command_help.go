package main

import "fmt"

func commandHelp(conf *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for k, v := range getCommands() {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	return nil
}
