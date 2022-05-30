package main

import "fmt"

func double(x int) (result int) {
	defer func() {
		// 可以获取到参数和返回值
		fmt.Printf("x=%v, result=%v\n", x, result)
	}()
	return x + x
}

func triple(x int) (result int) {
	defer func() {
		// 改变返回值
		result += 100
	}()
	return x + x
}

func main() {
	//double(2)
	result := triple(2)
	fmt.Println("result=", result)
}
