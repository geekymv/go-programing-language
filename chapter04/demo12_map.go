package main

import (
	"fmt"
)

func main() {

	var ages map[string]int
	fmt.Println("ages", ages == nil)    // true
	fmt.Println("len(ages)", len(ages)) // 0

	// 向 map 添加元素
	//ages["tom"] = 20 // panic

	kv := make(map[string]bool)
	kv["a"] = true
	kv["b"] = false

	// key nu c n
	fmt.Println("c = ", kv["c"])
}
