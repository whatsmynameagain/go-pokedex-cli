package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	mu       sync.Mutex
	lifetime time.Duration
}

// untested
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val, // it bothers me that these have the same name
	}
	c.entries[key] = entry
}

// untested
func (c *Cache) Get(key string) ([]byte, bool) {
	value, ok := c.entries[key]
	return value.val, ok
}

// untested
func (c *Cache) reapLoop(interval time.Duration) {

	reapTimer := time.NewTicker(interval)
	defer reapTimer.Stop()

	for {
		<-reapTimer.C
		//fmt.Println("Executing reapLoop()...")
		if len(c.entries) == 0 {
			//fmt.Println("Cache is empty, skipping...")
			return
		}

		c.mu.Lock()
		currentTime := time.Now()
		for key, entry := range c.entries {
			// if the difference between the current time and the entry's time
			// is greater than the cache's specified lifetime, delete it
			//fmt.Println("reapLoop: checking cache entries...")
			if currentTime.Sub(entry.createdAt) > c.lifetime {
				//fmt.Printf("reapLoop: deleting key\n")
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{
		entries:  make(map[string]cacheEntry),
		lifetime: interval,
	}

	go newCache.reapLoop(interval)
	return &newCache

}
