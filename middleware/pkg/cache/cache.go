// Package cache implements a cache. This cache is simple: a map with a mutex.
// There is no fancy expunge algorithm, it just randomly evicts elements when
// it gets full.
package cache

import "sync"

// onEvict?

// Cache is a cache with random eviction.
type Cache struct {
	items map[string]interface{}
	size  int

	sync.RWMutex
}

// New returns a new cache with the specified size.
func New(size int) *Cache { return &Cache{items: make(map[string]interface{}), size: size} }

// Add element indexed by key into the cache. Any existing
// element is overwritten
func (c *Cache) Add(key string, el interface{}) {
	l := c.Len()
	if l+1 > c.size {
		c.Evict()
	}

	// Now our locking.
	c.Lock()
	defer c.Unlock()
	c.items[key] = el
}

// Remove removes the element indexed by key from the cache.
func (c *Cache) Remove(key string) {
	c.Lock()
	defer c.Unlock()
	delete(c.items, key)
}

// Evict removes a random element from the cache.
func (c *Cache) Evict() {
	c.Lock()
	defer c.Unlock()

	key := ""
	for k := range c.items {
		key = k
		break
	}
	if key == "" {
		// empty cache
		return
	}
	delete(c.items, key)
}

// Get looks up the element indexed under key.
func (c *Cache) Get(key string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()

	el, found := c.items[key]
	return el, found
}

// Len returns the current length of the cache.
func (c *Cache) Len() int {
	c.RLock()
	defer c.RUnlock()
	return len(c.items)
}
