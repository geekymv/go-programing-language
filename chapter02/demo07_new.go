package main

import "fmt"

func main() {

	p := new(int)
	// *int 是一个 int 类型的指针
	fmt.Printf("p type:%T\n", p)

	*p = 100
	fmt.Println("p value:", *p)
}
