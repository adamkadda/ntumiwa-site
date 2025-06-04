package content

import (
	"sync"
	"time"
)

/*
	A combination of both TTL and TTI cache expiration
	strategies would be ideal, but TTI is not essential
	right now. Will consider implementing this later on.
*/

type Cache[T any] struct {
	data       *T
	mu         sync.RWMutex
	ttl        time.Duration
	lastParsed time.Time

	// tti        time.Duration
	// lastAccess time.Time
}

func (c *Cache[T]) Get() *T {
	c.mu.RLock()
	defer c.mu.RUnlock()

	now := time.Now()
	if c.lastParsed.IsZero() || now.Sub(c.lastParsed) > c.ttl {
		c.data = nil
	}

	return c.data
}

func (c *Cache[T]) Set(data *T) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data = data
	c.lastParsed = time.Now()
}
