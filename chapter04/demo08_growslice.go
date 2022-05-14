package main

import "fmt"

func main() {

	s := make([]int, 0, 1024)

	for i := 0; i < 1024; i++ {
		s = append(s, i)
	}
	fmt.Println(len(s), cap(s))

	s = append(s, 1)
	fmt.Println(len(s), cap(s))

}
