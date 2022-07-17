package main

import (
	"fmt"
	"sync"
)

var cache = make(map[int64]string)
var mu sync.RWMutex

func Get(key int64) string {
	mu.RLock()
	defer mu.RUnlock()
	return cache[key]
}

// g1  a:a1
// g2  a:a2
func Put(key int64, value string) string {
	// 先判断key是否存在
	mu.RLock()
	v, ok := cache[key]
	mu.RUnlock()
	if ok {
		return v
	}

	// 不存在
	// g1 获取到锁 g2 等待
	mu.Lock()
	// g1 执行完，释放锁，g2获取到锁，继续执行
	defer mu.Unlock()

	// 获取到锁，需要再次判断key是否存在（判断条件是否满足）
	v, ok = cache[key]
	if ok {
		return v
	}
	cache[key] = value
	return value
}

func main() {

	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			val := Put(1000, fmt.Sprintf("%d", i))
			fmt.Println(val)
		}(i)
	}
	wg.Wait()
	fmt.Println("val=", Get(1000))

}
