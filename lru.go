package lcache

import (
	"container/list"
	"errors"
)

// LRUCache implements a fixed size LRU cache.
type LRUCache struct {
	cache map[interface{}]*list.Element
	list  *list.List
	size  int
}

// element is used to store a value in LRUCache's list.
type element struct {
	key, value interface{}
}

// NewLRUCache returns new LRUCache.
func NewLRUCache(size int) (*LRUCache, error) {
	if size <= 0 {
		return nil, errors.New("size must be positive")
	}

	return &LRUCache{
		cache: make(map[interface{}]*list.Element),
		list:  list.New(),
		size:  size,
	}, nil
}

// Set sets a value to the cache.
func (lru *LRUCache) Set(key, value interface{}) {
	// if the cache contains provided key,
	// then move element belonged to the key to front of the list
	if item, ok := lru.cache[key]; ok {
		item.Value.(*element).value = value
		lru.list.MoveToFront(item)
		return
	}

	// remove oldest element
	if len(lru.cache) >= lru.size {
		back := lru.list.Back()
		delete(lru.cache, back.Value.(*element).key)
		lru.list.Remove(back)
	}

	// insert new element
	elem := &element{key: key, value: value}
	lru.cache[key] = lru.list.PushFront(elem)
}

// Get returns a key's value from the cache.
func (lru *LRUCache) Get(key interface{}) (interface{}, bool) {
	item, ok := lru.cache[key]
	if !ok {
		return nil, false
	}

	lru.list.MoveToFront(item)
	return item.Value.(*element).value, true
}

// Remove removes element from the cache by key.
// Returns true if the element was removed, false otherwise.
func (lru *LRUCache) Remove(key interface{}) bool {
	if item, ok := lru.cache[key]; ok {
		lru.list.Remove(item)
		delete(lru.cache, key)
		return true
	}

	return false
}

// Clear completely clears the cache.
func (lru *LRUCache) Clear() {
	for k := range lru.cache {
		delete(lru.cache, k)
	}

	lru.list.Init()
}

// Size returns number of the cache elements.
func (lru *LRUCache) Size() int {
	return lru.list.Len()
}
