package main

import (
	"fmt"
	"sync"
)

var balanceV1 int

// 存款
func DepositV1(amount int) {
	// 不是原子操作
	balanceV1 = balanceV1 + amount
}

// 查询余额
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
