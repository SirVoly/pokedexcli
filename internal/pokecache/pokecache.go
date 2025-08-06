package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	interval time.Duration
	entries  map[string]cacheEntry
	mu       *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		mu:       &sync.Mutex{},
		entries:  make(map[string]cacheEntry),
		interval: interval,
	}

	go func() {
		ticker := time.NewTicker(time.Millisecond)
		defer ticker.Stop()
		for {
			<-ticker.C
			c.reapLoop(interval)
		}
	}()

	return c
}

func (c Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = entry
}

func (c Cache) Get(key string) ([]byte, bool) {
	var val []byte
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	if !ok {
		return val, false
	}

	return entry.val, true
}

func (c Cache) reapLoop(duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, entry := range c.entries {
		if time.Since(entry.createdAt) > duration {
			delete(c.entries, key)
		}
	}
}
