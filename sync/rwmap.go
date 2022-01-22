package sync

import "sync"

type RWMap[K comparable, V any] struct {
	da map[K]V
	mu sync.RWMutex
}

var _ SyncMap[int, int] = &RWMap[int, int]{}

func (rwm *RWMap[K, V]) Get(k K) V {
	rwm.mu.RLock()
	v := rwm.da[k]
	rwm.mu.RUnlock()

	return v
}

func (rwm *RWMap[K, V]) Delete(k K) {
	rwm.mu.Lock()
	delete(rwm.da, k)
	rwm.mu.Unlock()
}

func (rwm *RWMap[K, V]) Store(k K, v V) {
	rwm.mu.Lock()
	rwm.da[k] = v
	rwm.mu.Unlock()
}

func (rwm *RWMap[K, V]) Load(k K) (V, bool) {
	rwm.mu.RLock()
	v, ok := rwm.da[k]
	rwm.mu.RUnlock()
	return v, ok
}

func (rwm *RWMap[K, V]) Range(f func(K, V) bool) {
	for k, v := range rwm.da {
		if !f(k, v) {
			break
		}
	}
}

func (rwm *RWMap[K, V]) LoadAndDelete(k K) (V, bool) {
	rwm.mu.Lock()
	v, ok := rwm.da[k]
	delete(rwm.da, k)
	rwm.mu.Unlock()

	return v, ok
}

func (rwm *RWMap[K, V]) LoadOrStore(k K, v V) (V, bool) {
	rwm.mu.Lock()
	v2, ok := rwm.da[k]
	if !ok {
		rwm.da[k] = v
		v2 = v
	}
	rwm.mu.Unlock()
	return v2, ok
}
