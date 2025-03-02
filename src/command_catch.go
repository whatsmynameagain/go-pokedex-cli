package main

import (
	"fmt"
	"math/rand"

	"github.com/whatsmynameagain/go-pokedex-cli/src/internal/pokeapi"
)

func commandCatch(conf *Config, args ...string) error {

	if len(args) == 1 {
		return fmt.Errorf("must enter pokemon name")
	} else if len(args) > 2 {
		return fmt.Errorf("must enter a single pokemon name")
	}

	toCatch := args[1]

	if _, ok := conf.caughtPokemon[toCatch]; ok {
		// to-do: automate the 'a/an'
		return fmt.Errorf("already caught a(n) %s", toCatch)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", toCatch)
	endPoint := "/pokemon/" + toCatch

	const minExp = 36
	const maxExp = 608
	// min 36 | max 608 according to
	//https://bulbapedia.bulbagarden.net/wiki/List_of_Pok%C3%A9mon_by_effort_value_yield_in_Generation_IX

	getURL := pokeapi.BaseURL + endPoint

	var pokemonInfo pokeapi.Pokemon
	pokemonInfo, err := conf.pokeapiClient.GetPokemon(getURL)
	if err != nil {
		return fmt.Errorf("error: %v\ntry again?", err)
	}

	pokemonBaseExp := pokemonInfo.BaseExperience
	normalizedExp := float64(pokemonBaseExp-minExp) / float64(maxExp-minExp)
	catchProbability := 0.9 - (0.8 * normalizedExp)
	roll := rand.Float64()

	if roll < catchProbability {
		fmt.Printf("%s was caught!\n", toCatch)
		// add to caught list
		conf.caughtPokemon[toCatch] = pokemonInfo
		fmt.Println("You may now inspect it with the inspect command")
	} else {
		fmt.Printf("%s escaped!\n", toCatch)
	}

	return nil
}
