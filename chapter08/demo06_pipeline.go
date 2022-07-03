package main

import (
	"fmt"
	"time"
)

func counter(out chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
		time.Sleep(1 * time.Second)
	}
	close(out)
}

func squarer(in <-chan int, out chan<- int) {
	for x := range in {
		fmt.Println("x=", x)
		out <- x * x
	}
	// close 只能用来关闭双向 channel 或只发送(send-only) channel
	close(out)
}

func print(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}

func main() {

	data := make(chan int)
	result := make(chan int)

	// counter 产生数据
	go counter(data)

	// squarer 计算平方
	go squarer(data, result)

	print(result)
}
