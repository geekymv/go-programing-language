package main

import "fmt"

// Runnable 抽象出一个接口, 让老师和学生自动实现这个接口
type Runnable interface {
	Run()
}

type Student struct {
}

func (s *Student) Run() {
	fmt.Println("student run")
}

type Teacher struct {
}

func (t *Teacher) Run() {
	fmt.Println("teacher run")
}

// Run 参数是一个接口类型
func Run(r Runnable) {
	r.Run()
}

func main() {

	s := &Student{}
	t := &Teacher{}

	Run(s)
	Run(t)
}
