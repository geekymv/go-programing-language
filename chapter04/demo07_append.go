package main

import "fmt"

func main() {

	var s []string // nil
	// nil 值的 slice 没有底层数组，slice 是空值可用的
	fmt.Println(len(s), cap(s))

	// append 函数返回一个新的 slice，容量足够的情况下返回的 slice 和原 slice 共用一个底层数组，
	// 容量不足，会分配一个新的底层数组
	s = append(s, "a")
	fmt.Println("s=", s)
	fmt.Println(len(s), cap(s))

	fmt.Println("-----------------")
	s = append(s, "b")
	fmt.Println(len(s), cap(s))

	fmt.Println("-----------------")
	s = append(s, "c")
	fmt.Println(len(s), cap(s))

	// len 表示slice元素个数，cap表示底层数组长度
	// slice 扩容到原cap的2倍（不一定是2倍）

	fmt.Println("-----------------")
	s = append(s, "d")
	fmt.Println(len(s), cap(s))

}
