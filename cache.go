package lcache

// Cache represents basic cache interface.
type Cache interface {
	// Set sets a cache element value,
	// if it's not presented - inserts new element.
	Set(key, value interface{})

	// Get returns a key's value from the cache.
	Get(key interface{}) (interface{}, bool)

	// Remove removes element from the cache.
	Remove(key interface{}) bool

	// Clears all cache elements.
	Clear()
}
