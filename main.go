package main

import (
	"time"

	"github.com/whatsmynameagain/go-pokedex-cli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 30*time.Second)
	configuration := &Config{
		pokeapiClient: &pokeClient,
	}
	start(configuration)
}
