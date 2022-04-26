package main

import "fmt"

func main() {

	var x int = 1
	// 声明一个int指针，它的值是变量x的地址
	var p *int = &x

	// 获取变量 x 的地址
	fmt.Printf("%p\n", &x)
	fmt.Printf("%v\n", p)

	// 变量 p 的类型
	fmt.Printf("%T\n", p)

	// * 访问存储在该地址中的值
	fmt.Println("获取变量的值", *p)

	// 更新变量的值
	*p = 2
	fmt.Println(x)

}
