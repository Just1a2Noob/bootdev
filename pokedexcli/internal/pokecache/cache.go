package pokecache

import "time"

func cache() map[string]cacheEntry {
	return map[string]cacheEntry{}
}

func NewCache(interval time.Duration) (cacheEntry, error) {
	return cacheEntry{}, nil
}

func (c *cacheEntry) Add(key string, val []byte) {
	c.val = val
}

func (c cacheEntry) Get(key string) ([]byte, bool) {
}

func (c cacheEntry) reapLoop() error {
	return nil
}
