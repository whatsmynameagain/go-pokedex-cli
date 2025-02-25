package pokeapi

/*
import (
	"github.com/whatsmynameagain/go-pokedex-cli/internal/pokeapi/pokecache"
)
*/

// check if the current url is cached
func (c *Client) CheckCache(getURL string) ([]byte, bool) {
	val, ok := c.cache.Get(getURL)
	if !ok {
		return []byte{}, ok
	}
	return val, false
}
