package main

import (
	"fmt"
	"sync"
)

var balance int

func Deposit(amount int) {
	balance = balance + amount
}

func Balance() int {
	return balance
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			Deposit(i)
		}(i)
	}
	wg.Wait()

	fmt.Println("balance=", Balance())

}
