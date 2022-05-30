package main

import "fmt"

func close(i int) {
	fmt.Println("close...", i)
}

func main() {

	fmt.Println("main start")
	// defer 最后执行，一般用于释放资源
	// 多个defer 倒序执行
	defer close(1)
	defer close(2)
	defer close(3)

	fmt.Println("main end")
}
