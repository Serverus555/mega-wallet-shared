package cache

import (
	"fmt"
	"time"
)

type Cache[K comparable, V any] interface {
	Put(key K, value V)
	Get(key K) (V, bool)
	Delete(key K)
}

type Type string

var (
	MemoryType Type = "memory"
	NoopType   Type = "noop"
	// redis
)

type Config struct {
	Type     Type
	Capacity int
	TTL      time.Duration
}

func New[K comparable, V any](c Config) (Cache[K, V], error) {
	switch c.Type {
	case MemoryType:
		return NewInMemory[K, V](c.Capacity, c.TTL), nil
	case NoopType:
		return NewNoop[K, V](), nil
	}
	return nil, fmt.Errorf("invalid cache type: %s", c.Type)
}
