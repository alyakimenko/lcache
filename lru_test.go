package lcache

import (
	"strconv"
	"testing"
)

func TestLRUCache_Set(t *testing.T) {
	cache, _ := NewLRUCache(5)

	for i := 0; i < 5; i++ {
		cache.Set(strconv.Itoa(i), i)
	}

	cache.Set("1", 2)

	item, ok := cache.Get("1")
	if !ok {
		t.Fatal("cache must contains '1' key")
	}
	if item.(int) != 2 {
		t.Fatal("item on '1' key must be 2")
	}

	cache.Set("6", 6)

	_, ok = cache.Get("0")
	if ok {
		t.Fatal("'0' key must not be presented in the cache")
	}
}

func TestLRUCache_Get(t *testing.T) {
	cache, _ := NewLRUCache(5)

	for i := 0; i < 5; i++ {
		key := strconv.Itoa(i)
		cache.Set(key, i)

		item, ok := cache.Get(key)
		if !ok {
			t.Fatalf("cache must contains %+v key", key)
		}

		if item != i {
			t.Fatalf("element under %+v key must be equal to %d", key, i)
		}
	}

	_, ok := cache.Get("6")
	if ok {
		t.Fatal("cache must not contains '6' key")
	}

	cache.Set("6", 6)
	_, ok = cache.Get("6")
	if !ok {
		t.Fatal("cache must contains '6' key")
	}

	_, ok = cache.Get("0")
	if ok {
		t.Fatal("cache must not contains '0' key")
	}
}

func TestLRUCache_Remove(t *testing.T) {
	cache, _ := NewLRUCache(5)

	for i := 0; i < 5; i++ {
		cache.Set(strconv.Itoa(i), i)
	}

	removed := cache.Remove("3")
	if !removed {
		t.Fatal("element on '3' key must be removed")
	}

	_, ok := cache.Get("3")
	if ok {
		t.Fatal("element on '3' key must not be presented in the cache")
	}

	if cache.Size() != 4 {
		t.Fatal("cache size must be 4")
	}
}

func TestLRUCache_Clear(t *testing.T) {
	cache, _ := NewLRUCache(5)

	for i := 0; i < 5; i++ {
		cache.Set(strconv.Itoa(i), i)
	}

	cache.Clear()

	if cache.Size() != 0 {
		t.Fatal("cache size must be 0")
	}
}
