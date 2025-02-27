package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/whatsmynameagain/go-pokedex-cli/internal/pokeapi/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout, lifetime time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: *pokecache.NewCache(lifetime),
	}
}

// get locations list
func (c *Client) GetLocationsList(getURL string) (ResponseLocations, error) {

	body, err := c.getPokeAPIData(getURL)

	if err != nil {
		return ResponseLocations{}, err
	}

	uResponse := ResponseLocations{}

	if err := json.Unmarshal(body, &uResponse); err != nil {
		return ResponseLocations{}, err
	}

	return uResponse, nil
}

// returns the body of the request after checking the cache
func (c *Client) getPokeAPIData(getURL string) ([]byte, error) {

	// check cache
	if cachedEntry, found := c.CheckCache(getURL); found {
		return cachedEntry, nil
	}

	// if no cache entry, run a get request
	res, err := http.Get(getURL)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return []byte{}, err
	}

	// update cache with new entry
	c.cache.Add(getURL, body)

	return body, nil
}

// check if the current url is cached
func (c *Client) CheckCache(getURL string) ([]byte, bool) {
	val, ok := c.cache.Get(getURL)
	if !ok {
		return []byte{}, ok
	}
	return val, false
}
