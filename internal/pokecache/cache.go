package pokecache

import (
	"sync"
	"time"
)

var mut sync.RWMutex

func NewCache(interval time.Duration) Cache {
	cache := Cache{data: make(map[string]cacheEntry)}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	mut.Lock()
	defer mut.Unlock()
	c.data[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	mut.RLock()
	defer mut.RUnlock()
	entry, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		for key, entry := range c.data {
			if time.Since(entry.createdAt) > interval {
				delete(c.data, key)
			}
		}
	}
}
