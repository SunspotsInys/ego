package maps

import "sync"

type RWMap[K comparable, V any] struct {
	data map[K]V
	mut  *sync.RWMutex
}

var _ SyncMap[int, int] = &RWMap[int, int]{}

func NewRWMap[K comparable, V any]() *RWMap[K, V] {
	return &RWMap[K, V]{
		data: map[K]V{},
		mut:  &sync.RWMutex{},
	}
}

func (rwm *RWMap[K, V]) Get(k K) V {
	rwm.mut.RLock()
	v := rwm.data[k]
	rwm.mut.RUnlock()

	return v
}

func (rwm *RWMap[K, V]) Delete(k K) {
	rwm.mut.Lock()
	delete(rwm.data, k)
	rwm.mut.Unlock()
}

func (rwm *RWMap[K, V]) Store(k K, v V) {
	rwm.mut.Lock()
	rwm.data[k] = v
	rwm.mut.Unlock()
}

func (rwm *RWMap[K, V]) Load(k K) (V, bool) {
	rwm.mut.RLock()
	v, ok := rwm.data[k]
	rwm.mut.RUnlock()
	return v, ok
}

func (rwm *RWMap[K, V]) Range(f func(K, V) bool) {
	rwm.mut.RLock()
	for k, v := range rwm.data {
		if !f(k, v) {
			break
		}
	}
	rwm.mut.RUnlock()
}

func (rwm *RWMap[K, V]) LoadAndDelete(k K) (V, bool) {
	rwm.mut.Lock()
	v, ok := rwm.data[k]
	delete(rwm.data, k)
	rwm.mut.Unlock()

	return v, ok
}

func (rwm *RWMap[K, V]) LoadOrStore(k K, v V) (V, bool) {
	rwm.mut.Lock()
	v2, ok := rwm.data[k]
	if !ok {
		rwm.data[k] = v
		v2 = v
	}
	rwm.mut.Unlock()
	return v2, ok
}
