package main

import (
	"fmt"
)

func commandInspect(conf *Config, args ...string) error {

	if len(args) == 1 {
		return fmt.Errorf("must enter pokemon name")
	} else if len(args) > 2 {
		return fmt.Errorf("must enter a single pokemon name")
	}

	// check if the pokemon to inspect has been caught
	toInspect := args[1]

	pokemon, ok := conf.caughtPokemon[toInspect]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	// display the info

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, v := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", v.Stat.Name, v.BaseStat)
	}
	fmt.Println("Types:")
	for _, v := range pokemon.Types {
		fmt.Printf(" - %v\n", v.Type.Name)
	}

	return nil
}
