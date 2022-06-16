package main

import "fmt"

type A interface {
	hello(s string) string

	world()
}

// B 类型B实现了A接口的所有方法，才能认为B实现了A接口
type B struct {
}

func (b *B) hello(s string) string {
	return "hello," + s
}

// world 方法的 receiver 是 *B
func (b *B) world() {
	fmt.Println("world")
}

func main() {

	var a A
	a = new(B)

	a.world()
}
