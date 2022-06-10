package cache

import "sync"

func New() Cache {
	return Cache{
		m: map[string]string{},
	}
}

type Cache struct {
	mutex sync.RWMutex
	m     map[string]string
}

func (c *Cache) Set(key, value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.m[key] = value
}

func (c *Cache) Get(key string) (value string) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.m[key]
}
