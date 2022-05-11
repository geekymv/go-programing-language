package main

import "fmt"

func main() {

	var a [3]int = [3]int{1, 2, 3}

	// [3]int 数组长度是数组类型的一部分
	fmt.Printf("%T\n", a)

	// [3]int [2]int 是不同类型
	var b = [2]int{1, 2}
	fmt.Printf("%T\n", b)

	c := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", c)

	// 给指定索引赋值
	r := [...]int{9: -1}
	fmt.Println(r)

	//i := [2]int{1, 2}
	//j := [3]int{1, 2}
	// fmt.Println(i == j)

}
