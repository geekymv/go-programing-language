package main

import (
	"fmt"
	"time"
)

// 什么时候关闭channel当：需要通知接收方所有数据已发送完毕的时候才需要关闭channel。
// 接收方如何判断channel 是否关闭。（ok标志，for range）
func main() {

	data := make(chan int)
	result := make(chan int)

	// counter 产生数据
	go func() {
		for x := 0; x < 10; x++ {
			data <- x
			time.Sleep(1 * time.Second)
		}
		// 数据发送完毕，关闭channel，表示数据发送完了
		close(data)
	}()

	// squarer 计算平方
	go func() {
		/*
			for {
				// 从已经关闭的channel中读取数据，读取完之后不会阻塞，获取到零值
				x, ok := <-data
				// ok 是 false 表示 data channel 已经关闭了
				if !ok {
					break
				}
				fmt.Println("x=", x)
				result <- x * x
			}
		*/
		// for range 读取channel上所有发送的值，直到接收完channel上最后一个值，退出循环。
		for x := range data {
			fmt.Println("x=", x)
			result <- x * x
		}
		close(result)
	}()

	/*for {
		v, ok := <-result
		if !ok {
			break
		}
		fmt.Println(v)
		fmt.Println(v)
	}*/

	for x := range result {
		fmt.Println(x)
	}
}
