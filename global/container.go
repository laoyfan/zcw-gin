package global

import "sync"

type Container struct {
	mu   *sync.RWMutex
	data map[string]interface{}
}

// NewContainer 初始化容器
func NewContainer() *Container {
	return &Container{
		mu:   new(sync.RWMutex),
		data: make(map[string]interface{}),
	}
}

// Search 查询容器对象
func (c *Container) Search(key string) (value interface{}, found bool) {
	c.mu.RLock()
	value, found = c.data[key]
	c.mu.RUnlock()
	return
}

// Set 向写入对象
func (c *Container) Set(key string, value interface{}) {
	c.mu.Lock()
	c.data[key] = value
	c.mu.Unlock()
}

// Get 获取容器对象
func (c *Container) Get(key string) (value interface{}) {
	c.mu.RLock()
	value, _ = c.data[key]
	c.mu.RUnlock()
	return
}

// Delete 从容器删除对象
func (c *Container) Delete(key string) (value interface{}) {
	c.mu.Lock()
	var ok bool
	if value, ok = c.data[key]; ok {
		delete(c.data, key)
	}
	c.mu.Unlock()
	return
}

// GetOrSetFunc 获取对象
func (c *Container) GetOrSetFunc(key string, f func() interface{}) interface{} {
	if v, ok := c.Search(key); !ok {
		return c.doSet(key, f)
	} else {
		return v
	}
}

// doSet 写入对象
func (c *Container) doSet(key string, value interface{}) interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()
	if v, ok := c.data[key]; ok {
		return v
	}
	if f, ok := value.(func() interface{}); ok {
		value = f()
	}
	if value != nil {
		c.data[key] = value
	}
	return value
}

// GetOrSetController 控制器对象
func (c *Container) GetOrSetController(key string, f func() interface{}) interface{} {
	return c.GetOrSetFunc("controller."+key, f)
}

// GetOrSetService 服务对象
func (c *Container) GetOrSetService(key string, f func() interface{}) interface{} {
	return c.GetOrSetFunc("service."+key, f)
}
