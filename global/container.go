package global

import "sync"

type Container struct {
	mu   *sync.RWMutex
	data map[string]interface{}
}

func NewContainer() *Container {
	return &Container{
		mu:   new(sync.RWMutex),
		data: make(map[string]interface{}),
	}
}

func (c *Container) Search(key string) (val interface{}, found bool) {
	c.mu.RLock()
	val, found = c.data[key]
	c.mu.RUnlock()
	return
}

func (c *Container) Set(key string, val interface{}) {
	c.mu.Lock()
	c.data[key] = val
	c.mu.Unlock()
}

func (c *Container) Get(key string) (val interface{}) {
	c.mu.RLock()
	val, _ = c.data[key]
	c.mu.RUnlock()
	return
}

func (c *Container) Delete(key string) (val interface{}) {
	c.mu.Lock()
	var ok bool
	if val, ok = c.data[key]; ok {
		delete(c.data, key)
	}
	c.mu.Unlock()
	return
}
