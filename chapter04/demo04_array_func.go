package main

import "fmt"

// 传递一个指针
func update2(p *[3]string) {
	p[0] = "d"
	fmt.Println("p=", *p)
}

func main() {

	a := [3]string{"a", "b", "c"}
	update2(&a)
	// a 没有改变
	fmt.Println("a=", a)
}
