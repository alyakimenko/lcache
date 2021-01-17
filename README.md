Simple, non-thread safe LRU cache implementation.

### Usage

```go
package main

import (
	"strconv"

	"github.com/alyakimenko/lcache"
)

const cacheSize = 5

func main() {
	cache, _ := lcache.NewLRUCache(cacheSize)
	
	// cache: {0, 1, 2, 3, 4}
	for i := 0; i < 5; i++ {
		cache.Set(strconv.Itoa(i), i)
	}
	
	// cache: {0, 1, 2, 4, 3}
	value, ok := cache.Get("3")
	if !ok {
		// not presented
	}
	
	// cache: {1, 2, 4, 3, 5}
	cache.Set("5")
	
	// cache: {1, 4, 3, 5}
	removed := cache.Remove("2")
	
	// cache: {}
	cache.Clear()
}
```