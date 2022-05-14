package main

import "fmt"

func main() {

	// len = cap = 3
	s1 := make([]int, 3)

	// len = 3 cap = 6
	s2 := make([]int, 3, 6)

	fmt.Println(s1, s2)
}
