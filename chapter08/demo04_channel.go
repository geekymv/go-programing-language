package main

import (
	"fmt"
	"time"
)

func run(ch chan struct{}) {
	for i := 0; i < 10; i++ {
		fmt.Println("i=", i)
		time.Sleep(1 * time.Second)
	}
	ch <- struct{}{}
}

func main() {
	fmt.Println("main start")
	// 无缓冲channel 发送和接收数据会阻塞
	ch := make(chan struct{})
	go run(ch)

	// 其他业务操作

	fmt.Println("wait...")
	<-ch
	fmt.Println("main exit")

}
