package main

import (
	"fmt"

	"github.com/whatsmynameagain/go-pokedex-cli/internal/pokeapi"
)

func commandExplore(conf *Config, args ...string) error {

	if len(args) == 1 {
		return fmt.Errorf("must enter location name")
	} else if len(args) > 2 {
		return fmt.Errorf("must enter a single location name")
	}

	endPoint := "/location-area/" + args[1] + "/"

	getURL := pokeapi.BaseURL + endPoint

	fmt.Printf("Exploring %s...\n", args[1])

	locationPokemonResp, err := conf.pokeapiClient.GetLocationPokemon(getURL)
	if err != nil {
		return fmt.Errorf("error: %v\ntry again?", err)
	}

	fmt.Println("Found Pokemon:")

	for _, encounter := range locationPokemonResp.PokemonEncounters {
		fmt.Printf("- %v\n", encounter.Pokemon.Name)
	}
	// PokemonEncounters -> Pokemon -> Name
	return nil
}
