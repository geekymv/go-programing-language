package main

import "fmt"

type Point struct {
	X int
	Y int
}

// IncrAge1 按值传递
func IncrAge1(u User) {
	u.Age += 1
}

func IncrAge2(u *User) {
	u.Age += 1
}

func main() {

	p := Point{1, 2}
	fmt.Printf("%#v\n", p)

	// 常用的形式
	p2 := Point{
		X: 11,
		Y: 22,
	}
	fmt.Printf("%#v\n", p2)

	fmt.Println("-------------------------")

	u := User{
		ID:   1,
		Name: "jack",
		Age:  20,
	}

	IncrAge2(&u)
	fmt.Printf("%#v\n", u)

}
