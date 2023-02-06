package main

import (
	"fmt"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func sendMsg(u User) {
	fmt.Println("send msg to ", u.Name)
}

func main() {

	users := []*User{
		{
			ID:   1,
			Name: "tom",
			Age:  20,
		},
		{
			ID:   2,
			Name: "jack",
			Age:  21,
		},
		{
			ID:   3,
			Name: "lily",
			Age:  22,
		},
	}

	for _, u := range users {
		u.Age += 1
	}

	for _, u := range users {
		fmt.Printf("%v, %d\n", u.Name, u.Age)
	}
}
