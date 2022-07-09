package main

import (
	"fmt"
	"os"
	"time"
)

var done = make(chan int)

func cancelled() bool {
	select {
	case v := <-done: // 从已经关闭的channel中读取数据，读取完之后不会阻塞，获取到零值
		fmt.Println("done=", v)
		return true
	default:
		return false
	}
}

func main() {

	go func() {
		// 从标准输入中读取一个字符
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	for {
		time.Sleep(1 * time.Second)
		flag := cancelled()
		fmt.Println("flag=", flag)
	}

}
