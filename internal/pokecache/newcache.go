package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{
		mapCache: make(map[string]cacheEntry),
		mu:       sync.Mutex{},
	}
	go newCache.reapLoop(interval)
	return newCache
}
