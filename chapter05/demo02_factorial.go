package main

import "fmt"

// n! = n * (n-1)!
// 0! = 1
func fact(n int) int {

	if n == 0 {
		return 1
	}

	// 递归调用
	return n * fact(n-1)
}

func main() {
	// 5! = 5 * 4 * 3 * 2 * 1
	r := fact(5)
	fmt.Println("r=", r)
}
