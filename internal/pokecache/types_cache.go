package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val}
}

func (c *Cache) Get(key string, val []byte) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	if !ok {
		return []byte{}, false
	}

	return entry.val, true
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(creation time.Time, data []byte) cacheEntry {
	return cacheEntry{
		createdAt: creation,
		val:       data,
	}
}
