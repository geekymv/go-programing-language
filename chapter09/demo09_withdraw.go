package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	// 互斥量本身和被保护对象都没有导出
	// RWMutex 多用于读多写少的场景。
	muV8      sync.RWMutex
	balanceV8 int
)

func DepositV8(amount int) {
	// 写锁互斥
	muV8.Lock()
	defer muV8.Unlock()
	depositV8(amount)
}

func BalanceV8() int {
	// 读锁可以并发执行（共享锁）
	muV8.RLock()
	defer muV8.RUnlock()
	return balanceV8
}

// 封装一个不可导出函数，供其他函数调用，没有添加锁
func depositV8(amount int) {
	balanceV8 += amount
}

func WithDrawV8(amount int) int {
	// 只获取一次锁
	muV8.Lock()
	defer muV8.Unlock()

	if balanceV8 < amount {
		return balanceV8
	}

	time.Sleep(100 * time.Millisecond)
	depositV8(-amount)
	return balanceV8
}

func main() {
	balanceV8 = 10
	var wg sync.WaitGroup
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			res := WithDrawV8(i)
			if res < 0 {
				fmt.Println(res)
			}
		}(i)
	}
	wg.Wait()

	fmt.Println("balance", balanceV8)
}
