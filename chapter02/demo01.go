package main

import "fmt"

// 包级别的变量
var i int = 1

// PI 包级别的常量
const PI = 3.14

func main() {
	fmt.Println("i=", i)
}

func say(name string) (int, string) {
	fmt.Println("i=", i)
	// 多返回值
	return i, "hello, " + name
}
