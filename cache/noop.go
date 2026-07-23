package cache

type Noop[K comparable, V any] struct {
}

func NewNoop[K comparable, V any]() *Noop[K, V] {
	return &Noop[K, V]{}
}

func (c *Noop[K, V]) Put(key K, value V) {
}

func (c *Noop[K, V]) Get(key K) (V, bool) {
	var zero V
	return zero, false
}

func (c *Noop[K, V]) Delete(key K) {
}
