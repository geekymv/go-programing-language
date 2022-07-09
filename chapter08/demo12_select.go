package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	abort := make(chan struct{})
	go func() {
		// 从标准输入中读取一个字符
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	tick := time.Tick(1 * time.Second)
	for count := 10; count > 0; count-- {
		// select 在多个 channel 上 多路复用
		select {
		case <-tick:
			fmt.Println(count)
		case <-abort:
			fmt.Println("abort")
			return
		}
	}

	fmt.Println("launch")
}
