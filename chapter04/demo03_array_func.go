package main

import "fmt"

// 值传递，创建一个副本
func update(p [3]string) {
	p[0] = "d"
	fmt.Println("p=", p)
}

func main() {

	a := [3]string{"a", "b", "c"}
	update(a)
	// a 没有改变
	fmt.Println("a=", a)
}
