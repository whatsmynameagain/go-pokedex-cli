package pokeapi

import (
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
