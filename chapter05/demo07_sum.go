package main

import "fmt"

func sum(vals ...int) int {
	fmt.Printf("vals type:%T\n", vals) // []int
	total := 0
	for _, v := range vals {
		total += v
	}
	return total
}

func main() {

	//fmt.Println("sum=", sum())
	//fmt.Println("sum=", sum(1))
	//fmt.Println("sum=", sum(1, 2, 3, 4))
	values := []int{1, 2, 3, 4, 5}
	fmt.Println("sum=", sum(values...))

}
