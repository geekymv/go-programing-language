package main

import "fmt"

func main() {

	// 栈 先进后出
	stack := make([]int, 0, 10)

	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)

	top := stack[len(stack)-1]
	fmt.Println("top", top)

	pop := stack[0 : len(stack)-1]
	fmt.Println("pop", pop)
}
