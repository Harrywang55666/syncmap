package syncmap

import "sync"

type SyncMap[K comparable, V any] struct{ sm sync.Map }

func NewSyncMap[K comparable, V any]() *SyncMap[K, V] {
	return new(SyncMap[K, V])
}

// CompareAndDelete deletes the entry for key if its value is equal to old.
// The old value must be of a comparable type.
//
// If there is no current value for key in the map, CompareAndDelete
// returns false (even if the old value is the nil interface value).
func (sm *SyncMap[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	return sm.sm.CompareAndDelete(key, old)
}

// Delete deletes the value for a key.
func (sm *SyncMap[K, V]) Delete(key K) {
	sm.sm.Delete(key)
}

// Load returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.
func (sm *SyncMap[K, V]) Load(key K) (value V, ok bool) {
	v, ok := sm.sm.Load(key)
	if v == nil {
		return
	}
	return v.(V), ok
}

// LoadAndDelete deletes the value for a key, returning the previous value if any.
// The loaded result reports whether the key was present.
func (sm *SyncMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	v, ok := sm.sm.LoadAndDelete(key)
	if v == nil {
		return
	}
	return v.(V), ok
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (sm *SyncMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	v, ok := sm.sm.LoadOrStore(key, value)
	if v == nil {
		return
	}
	return v.(V), ok
}

// CompareAndSwap swaps the old and new values for key
// if the value stored in the map is equal to old.
// The old value must be of a comparable type.
func (sm *SyncMap[K, V]) CompareAndSwap(key K, old, new V) bool {
	return sm.sm.CompareAndSwap(key, old, new)
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
//
// Range does not necessarily correspond to any consistent snapshot of the Map's
// contents: no key will be visited more than once, but if the value for any key
// is stored or deleted concurrently (including by f), Range may reflect any
// mapping for that key from any point during the Range call. Range does not
// block other methods on the receiver; even f itself may call any method on m.
//
// Range may be O(N) with the number of elements in the map even if f returns
// false after a constant number of calls.
func (sm *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	sm.sm.Range(func(key, value any) bool { return f(key.(K), value.(V)) })
}

// Store sets the value for a key.
func (sm *SyncMap[K, V]) Store(key K, value V) {
	sm.sm.Store(key, value)
}

// Swap swaps the value for a key and returns the previous value if any.
// The loaded result reports whether the key was present.
func (sm *SyncMap[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	v, ok := sm.sm.Swap(key, value)
	if v == nil {
		return
	}
	return v.(V), ok
}
