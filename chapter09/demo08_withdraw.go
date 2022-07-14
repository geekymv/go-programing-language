package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	muV7      sync.Mutex
	balanceV7 int
)

func DepositV7(amount int) {
	muV7.Lock()
	defer muV7.Unlock()
	depositV7(amount)
}

func BalanceV7() int {
	muV7.Lock()
	defer muV7.Unlock()
	return balanceV7
}

func depositV7(amount int) {
	balanceV7 += amount
}

func WithDrawV7(amount int) int {
	muV7.Lock()
	defer muV7.Unlock()

	if balanceV7 < amount {
		return balanceV7
	}
	time.Sleep(10 * time.Millisecond)
	depositV7(-amount)
	return balanceV7
}

func main() {
	balanceV7 = 10
	var wg sync.WaitGroup
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			res := WithDrawV7(i)
			if res < 0 {
				fmt.Println(res)
			}
		}(i)
	}
	wg.Wait()

	fmt.Println("balance", balanceV7)
}
