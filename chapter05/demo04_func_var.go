package main

import (
	"errors"
	"fmt"
	"os"
)

func negative(n int) int {
	return -n
}

func handle(i, j int, f func(int, int) (int, error)) (int, error) {
	if f == nil {
		return 0, errors.New("func f is nil")
	}
	return f(i, j)
}

func addOp(i, j int) (int, error) {
	return i + j, nil
}

func subOp(i, j int) (int, error) {
	return i - j, nil
}

func divOP(i, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("除数不能为0")
	}
	return i / j, nil
}

func main() {

	// f 函数变量
	f := negative
	fmt.Printf("%T\n", f)

	res := f(3)
	fmt.Println("res=", res)

	// 函数作为参数传递
	add, err := handle(3, 2, addOp)
	if err != nil {
		fmt.Printf("handle err:%v\n", err)
		os.Exit(1)
	}
	fmt.Println("add=", add)

	sub, err := handle(3, 2, subOp)
	if err != nil {
		fmt.Printf("handle err:%v\n", err)
	}
	fmt.Println("sub=", sub)
}
