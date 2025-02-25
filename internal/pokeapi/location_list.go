package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//get locations list

func (c Client) GetLocationsList(getURL string) (ResponseLocations, error) {

	res, err := http.Get(getURL)
	if err != nil {
		return ResponseLocations{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	// not sure if this is needed, copied it from
	// https://pkg.go.dev/net/http#example-Get
	// will remove if redundant
	if res.StatusCode > 299 {
		return ResponseLocations{}, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}

	if err != nil {
		return ResponseLocations{}, err
	}

	uResponse := ResponseLocations{}

	if err := json.Unmarshal(body, &uResponse); err != nil {
		return ResponseLocations{}, err
	}

	return uResponse, nil
}
