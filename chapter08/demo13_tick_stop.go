package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.NewTicker(1 * time.Second)

	go func() {
		for {
			select {
			case <-tick.C:
				fmt.Println("do something")
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("do some other thing")

	// 指定退出机制
	tick.Stop()
	fmt.Println("stop tick")

	time.Sleep(5 * time.Second)
}
