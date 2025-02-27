package main

import (
	"fmt"

	"github.com/whatsmynameagain/go-pokedex-cli/internal/pokeapi"
)

func commandMapF(conf *Config) error {
	return commandMap(conf, true)
}

func commandMapB(conf *Config) error {
	if conf.Previous == nil {
		return fmt.Errorf("you're on the first page")
	}
	return commandMap(conf, false)
}

func commandMap(conf *Config, forward bool) error {

	var getURL string
	// if moving forwards or backwards, check if Next or Previous are empty
	// if so, use the default
	switch forward {
	case true:
		if conf.Next == nil {
			getURL = pokeapi.BaseURL + "/location-area/"
		} else {
			getURL = *conf.Next
		}
	case false:
		if conf.Previous == nil {
			getURL = pokeapi.BaseURL + "/location-area/"
		} else {
			getURL = *conf.Previous
		}
	}

	locationsResp, err := conf.pokeapiClient.GetLocationsList(getURL)
	if err != nil {
		fmt.Println("Response error: ")
		fmt.Printf("%v\n", err)
		fmt.Println("Try again?")
		return fmt.Errorf("response error: \n%v\ntry again?", err)
	}

	conf.Next = locationsResp.Next
	conf.Previous = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Printf("%v\n", location.Name)
	}

	return nil
}
