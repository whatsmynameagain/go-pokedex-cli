package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(conf *Config) error {

	locations, err := getLocationsList(conf)
	if err != nil {
		fmt.Println("Response error: ")
		fmt.Printf("%v\n", err)
		fmt.Println("Try again?")
		return fmt.Errorf("Response error: \n%v\nTry again?", err)
	}
	for _, location := range locations {
		fmt.Printf("%v\n", location)
	}

	return nil
}

func getLocationsList(conf *Config) ([]string, error) {

	// use the base url if there's no Next
	var getURL string
	if conf.Next == "" {
		getURL = "https://pokeapi.co/api/v2/location-area/"
	} else {
		getURL = conf.Next
	}

	res, err := http.Get(getURL)
	if err != nil {
		return []string{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	// not sure if this is needed, copied it from
	// https://pkg.go.dev/net/http#example-Get
	// will remove if redundant
	if res.StatusCode > 299 {
		return []string{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}

	if err != nil {
		return []string{}, err
	}

	//fmt.Println(string(body))

	var uResponse Response

	if err := json.Unmarshal(body, &uResponse); err != nil {
		return []string{}, err
	}

	var out []string
	for _, result := range uResponse.Results {
		out = append(out, result.Name)
	}

	// update config
	conf.Next = uResponse.Next
	conf.Previous = uResponse.Previous
	return out, nil
}

// actual locations
// generated with JSON-to-GO
type Location struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

// response from GET
type Response struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
