package main

import (
	"time"

	"github.com/whatsmynameagain/go-pokedex-cli/src/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Second)
	caughtList := make(map[string]pokeapi.Pokemon)
	configuration := &Config{
		pokeapiClient: &pokeClient,
		caughtPokemon: caughtList,
	}
	start(configuration)
}
