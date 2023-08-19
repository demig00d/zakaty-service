package ttlcache

import (
	"sync"
	"time"
)

type TTLCache[T any] struct {
	value      T
	lastAccess int64
	expired    bool
	l          sync.Mutex
}

func NewTTLCache[T any](maxTTL int) (m *TTLCache[T]) {
	cache := TTLCache[T]{expired: true}

	go func() {
		for now := range time.Tick(time.Second) {
			m.l.Lock()
			if now.Unix()-cache.lastAccess > int64(maxTTL) {
				cache.expired = true
			}
			m.l.Unlock()
		}
	}()

	return &cache

}

func (cache *TTLCache[T]) Put(t T) {
	cache.l.Lock()
	cache.value = t
	cache.expired = false
	cache.lastAccess = time.Now().Unix()
	cache.l.Unlock()
}

func (cache *TTLCache[T]) Get() (t T, expired bool) {
	cache.l.Lock()
	t = cache.value
	expired = cache.expired
	cache.lastAccess = time.Now().Unix()
	cache.l.Unlock()
	return
}
