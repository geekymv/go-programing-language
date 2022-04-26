package main

import "fmt"

func main() {

	var x int = 100

	// 声明 *int 指针变量，只能存储 int 类型的内存地址
	var p *int = &x

	fmt.Println("x addr", &x)
	fmt.Println("p value", p)

	fmt.Println("x value", x)
	// * 在指针变量前面使用是取值操作，用于获取指针变量p指向的内存地址的值
	fmt.Println("x value", *p)

	// 通过指针改变变量x的值
	*p = 500
	fmt.Println("x value", x)

}
