package main

import "fmt"

func main() {
	// 声明一个变量
	var name string = "zhangsan"

	// 可以省略类型
	var age = 1

	// 可以省略表达式，初始值对应类型的零值
	var no int

	var s string

	fmt.Println("name=", name)
	fmt.Println("age=", age)
	fmt.Println("no=", no)
	fmt.Println("s=", s)

}
