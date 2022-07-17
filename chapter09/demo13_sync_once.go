package main

import (
	"fmt"
	"sync"
)

// sync.Once 只执行一次，零值可用。
var once sync.Once

func InitData() {
	// 只执行一次
	once.Do(func() {
		fmt.Println("init data...")
	})
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			InitData()
		}()
	}

	wg.Wait()
}
