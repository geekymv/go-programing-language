package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	muV6      sync.Mutex
	balanceV6 int
)

func DepositV6(amount int) {
	muV6.Lock()
	defer muV6.Unlock()
	balanceV6 += amount
}

func BalanceV6() int {
	muV6.Lock()
	defer muV6.Unlock()
	return balanceV6
}

func WithDrawV6(amount int) int {
	muV6.Lock()
	defer muV6.Unlock()

	if BalanceV6() < amount {
		return balanceV6
	}
	time.Sleep(10 * time.Millisecond)
	DepositV6(-amount)
	return balanceV6
}

func main() {
	balanceV6 = 10
	var wg sync.WaitGroup
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			res := WithDrawV6(i)
			if res < 0 {
				fmt.Println(res)
			}
		}(i)
	}
	wg.Wait()

	fmt.Println("exit")
}
