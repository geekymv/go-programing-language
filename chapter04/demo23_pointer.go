package main

import "fmt"

func main() {

	var a int = 10
	var pa *int = &a

	fmt.Printf("pa:%x\n", pa) // 16进制
	fmt.Printf("pa:%v\n", pa)

	fmt.Printf("a:%v\n", a)    // 10
	fmt.Printf("a:%b\n", a)    // 10 的二进制表示 -> 1010
	fmt.Printf("pa:%v\n", *pa) // 10
}
