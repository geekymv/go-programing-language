package main

import (
	"fmt"
	"sync"
)

var (
	// sync.Mutex 零值可用，默认是无锁，不可重入锁
	// 类似 Java ReentrantLock 可重入锁（可以再次获取锁）
	muV4      sync.Mutex
	balanceV4 int
)

func DepositV4(amount int) {
	// 获取锁，获取不到锁，会等待
	muV4.Lock()
	// 释放锁
	defer muV4.Unlock()
	balanceV4 += amount
}

func BalanceV4() int {
	muV4.Lock()
	defer muV4.Unlock()
	return balanceV4
}

func main() {

	var wg sync.WaitGroup
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			DepositV4(i)
		}(i)
	}
	wg.Wait()

	fmt.Println("balance=", BalanceV4())

}
