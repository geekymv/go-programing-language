package main

import (
	"fmt"
	"sync"
)

var (
	muV4      sync.Mutex
	balanceV4 int
)

func DepositV4(amount int) {
	muV4.Lock()
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
