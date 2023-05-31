package main

import (
	"fmt"
	"time"
)

func main() {

	hello(1)

	fmt.Println("111111")

	time.Sleep(time.Second * 20)

}

func hello(i int) int {
	var cnt int
	ticker := time.NewTicker(time.Second * 2)
	for {
		select {
		case <-ticker.C:
			fmt.Println("111")
			cnt++
			if cnt > 3 {
				fmt.Println(">>>")
				return 1
			}
		}
	}
}
