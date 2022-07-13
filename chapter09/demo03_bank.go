package main

import (
	"fmt"
	"sync"
)

var deposits = make(chan int)
var balances = make(chan int)

func DepositV2(amount int) {
	deposits <- amount
}

func BalanceV2() int {
	return <-balances
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}

func main() {

	var wg sync.WaitGroup
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			DepositV2(i)
		}(i)
	}
	wg.Wait()

	fmt.Println("balance=", BalanceV2())

}
