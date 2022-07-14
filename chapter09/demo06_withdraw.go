package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	muV5      sync.Mutex
	balanceV5 int
)

func DepositV5(amount int) {
	muV5.Lock()
	defer muV5.Unlock()
	balanceV5 += amount
}

func BalanceV5() int {
	muV5.Lock()
	defer muV5.Unlock()
	return balanceV5
}

// 	取款操作（不是原子操作）
func WithDrawV5(amount int) int {
	if BalanceV5() < amount {
		return balanceV5
	}
	// 模拟业务逻辑操作
	time.Sleep(10 * time.Millisecond)
	DepositV5(-amount)
	return balanceV5
}

func main() {
	balanceV5 = 10
	var wg sync.WaitGroup
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			res := WithDrawV5(i)
			if res < 0 {
				fmt.Println(res)
			}
		}(i)
	}
	wg.Wait()

	fmt.Println("exit")
}
