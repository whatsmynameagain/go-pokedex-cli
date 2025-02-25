package pokecache

import (
	"fmt"
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
func (c *Cache) reapLoop() {
	fmt.Println("Executing reapLoop()...")
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.entries {
		currentTime := time.Now()
		// if the difference between the current time and the entry's time
		// is greater than the cache's specified lifetime, delete it
		if diff := currentTime.Sub(entry.createdAt); diff > c.lifetime {
			delete(c.entries, key)
		}
	}
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{
		entries: make(map[string]cacheEntry),
		// no need to touch the mutex field
		// defer unlock
		lifetime: interval,
	}
	reapTimer := time.NewTicker(interval)
	defer reapTimer.Stop()
	go func() {
		for {
			<-reapTimer.C
			newCache.reapLoop()
		}
	}()
	return &newCache
}
