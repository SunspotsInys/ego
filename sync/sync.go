package sync

type SyncMap[K comparable, V any] interface {
	Get(K) V
	Delete(K)
	Store(K, V)
	Load(K) (V, bool)
	Range(func(K, V) bool)
	LoadAndDelete(K) (V, bool)
	LoadOrStore(K, V) (V, bool)
}