package main

import (
	"fmt"
	"runtime"
)

/*
Go调度模型GPM
G:Goroutine
P:逻辑Processor
M:OS Thread
https://deepu.tech/memory-management-in-golang/
https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html

栈帧
Go调度器
用户态、内核态

超线程
*/

func main() {

	/*
		cpu := runtime.NumCPU()
		fmt.Println(cpu)
	*/
	runtime.GOMAXPROCS(1)
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}

}
