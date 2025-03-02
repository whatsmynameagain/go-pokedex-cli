package main

import (
	"fmt"
)

func commandPokedex(conf *Config, args ...string) error {

	if len(conf.caughtPokemon) == 0 {
		fmt.Println("Your Pokedex is empty!")
		return nil
	}

	fmt.Println("Your pokedex:")
	for _, v := range conf.caughtPokemon {
		fmt.Printf("  - %s\n", v.Name)
	}

	return nil
}
