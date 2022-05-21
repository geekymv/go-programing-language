package main

import "fmt"

// i j 形参
func add(i, j int) (z int) {
	fmt.Println("z = ", z)
	z = i + j
	return
}

func main() {

	// 1 2 实参
	z := add(1, 2)
	fmt.Println("z = ", z)
}
