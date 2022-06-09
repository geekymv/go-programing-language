package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// Say 声明一个方法
func (u User) Say(content string) {
	fmt.Println(u.Name+" say:", content)
}

// Run 指针 receiver
func (u *User) Run() {
	fmt.Println(u.Name + " run...")
}

func main() {

	u1 := User{
		Name: "tom",
		Age:  20,
	}

	u2 := User{
		Name: "jack",
		Age:  20,
	}

	u3 := &User{
		Name: "kitty",
		Age:  18,
	}

	// 1.形参receiver 和实参 receiver 是同一类型
	u1.Say("hello")
	u2.Say("go")
	u3.Run()

	// 2.形参receiver *T类型 和实参 receiver 是T类型
	u1.Run() // (&u1).Run()
	u2.Run()

	// 3.形参receiver T类型 和实参 receiver 是*T类型
	u3.Say("hi") // 编译器隐式转换 (*u3).Run()
}
