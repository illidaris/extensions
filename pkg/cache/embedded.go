package cache

import (
	"sync"
	"time"
)

var embeddedCacheCore sync.Map

type EmbeddedCache struct {
}

func NewEmbeddedCache() *EmbeddedCache {
	return &EmbeddedCache{}
}

func (c *EmbeddedCache) Get(key string) interface{} {
	v, ok := embeddedCacheCore.Load(key)
	if ok {
		return v
	}
	return nil
}

func (c *EmbeddedCache) Set(key string, val interface{}, timeout time.Duration) error {
	embeddedCacheCore.Store(key, val)
	return nil
}

func (c *EmbeddedCache) SetNX(key string, val interface{}, timeout time.Duration) (bool, error) {
	embeddedCacheCore.Store(key, val)
	return true, nil
}

func (c *EmbeddedCache) IsExist(key string) bool {
	_, ok := embeddedCacheCore.Load(key)
	return ok
}

func (c *EmbeddedCache) Delete(key string) error {
	embeddedCacheCore.Delete(key)
	return nil
}
