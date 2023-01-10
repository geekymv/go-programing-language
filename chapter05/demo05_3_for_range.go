package main

import "fmt"

func main() {
	// 切片s
	s := []int{1, 1, 1}
	// 数组c
	c := [3]int{}
	for i, v := range s {
		if i == 0 {
			s[1] = 2
			s[2] = 2
		}
		c[i] = v
	}
	fmt.Println("s=", s)
	fmt.Println("c=", c)
}
