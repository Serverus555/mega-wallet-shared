package cache

import (
	"time"
)

type InMemory[K comparable, V any] struct {
	cache *otter.Cache[K, V]
}

func NewInMemory[K comparable, V any](capacity int, ttl time.Duration) *InMemory[K, V] {
	return &InMemory[K, V]{
		cache: otter.Must(&otter.Options[K, V]{
			MaximumSize:      capacity,
			ExpiryCalculator: otter.ExpiryWriting[K, V](ttl),
		}),
	}
}

func (c *InMemory[K, V]) Put(key K, value V) {
	c.cache.Set(key, value)
}

func (c *InMemory[K, V]) Get(key K) (V, bool) {
	return c.cache.GetIfPresent(key)
}

func (c *InMemory[K, V]) Delete(key K) {
	c.cache.Invalidate(key)
}
