package main

import (
	"flag"
	"fmt"
)

func main() {

	//args := os.Args[1:]
	//fmt.Println(args)

	// 返回 *bool 是一个 bool 类型的指针
	n := flag.Bool("n", false, "new line")

	// 解析命令行参数，需要在定义之后访问之前执行 Parse()
	flag.Parse()

	fmt.Println("n value", *n)

	// 获取命令行参数 等价 os.Args[1:] 不包括flag声明的命令行参数
	flag.Args()

}
