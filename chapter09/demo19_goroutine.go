package main

import (
	"fmt"
)

/*
Go调度模型GPM
G:Goroutine
P:逻辑Processor
M:OS Thread
https://deepu.tech/memory-management-in-golang/
https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html
*/

func main() {
	//runtime.GOMAXPROCS(1)
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}

}
