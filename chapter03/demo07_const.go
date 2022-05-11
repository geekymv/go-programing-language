package main

import (
	"fmt"
	"time"
)

const (
	a = 1
	b
	c = 2
	d
)

type Flags uint

const (
	FlagUp        Flags = 1 << iota // iota=0
	FlagBroadcast                   // 1 << iota // iota=1
	FlagLoopback                    // 1 << iota // iota=2
)

func main() {

	const pi = 3.14

	/*i := 100 / 0
	fmt.Println(i)*/

	/*	str := "hello"
		s := str[6]
		fmt.Println(s)*/

	fmt.Println(a, b, c, d)

	fmt.Println(int(time.Sunday))
	fmt.Println(int(time.Monday))

	fmt.Println("----------------")
	fmt.Println(FlagUp)
	fmt.Println(FlagBroadcast)
	fmt.Println(FlagLoopback)

}
