package main

import (
	"fmt"
	"sync"
)

var balanceV1 int

func DepositV1(amount int) {
	balanceV1 = balanceV1 + amount
}

func BalanceV1() int {
	return balanceV1
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			DepositV1(i)
		}(i)
	}
	wg.Wait()

	fmt.Println("balance=", BalanceV1())

}
