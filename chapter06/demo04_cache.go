package main

import (
	"fmt"
	"sync"
)

// Cache 封装一个缓存
type Cache struct {
	sync.Mutex
	mapping map[string]string
}

func (c Cache) Lookup(key string) string {
	c.Lock()
	defer c.Unlock()

	return c.mapping[key]
}

func (c Cache) Add(key, value string) {
	c.Lock()
	defer c.Unlock()

	c.mapping[key] = value
}

func main() {

	cache := Cache{
		mapping: make(map[string]string),
	}
	cache.Add("hello", "gogo")

	value := cache.Lookup("hello")
	fmt.Println("value=", value)

}
