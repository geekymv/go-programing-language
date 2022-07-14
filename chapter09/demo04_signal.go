package main

import (
	"fmt"
	"sync"
)

var (
	signal    = make(chan struct{}, 1)
	balanceV3 int
)

func DepositV3(amount int) {
	signal <- struct{}{}
	defer func() {
		<-signal
	}()
	balanceV3 += amount
}

func BalanceV3() int {
	signal <- struct{}{}
	defer func() {
		<-signal
	}()
	return balanceV3
}

func main() {

	var wg sync.WaitGroup
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			DepositV3(i)
		}(i)
	}
	wg.Wait()

	fmt.Println("balance=", BalanceV3())

}
