package main

import "fmt"

func main() {

	var a [3]int
	fmt.Println("len=", len(a))
	fmt.Println("a[0]=", a[0])
	fmt.Println("a[len(a)-1]=", a[len(a)-1])

	for i, v := range a {
		fmt.Printf("i=%v, v=%v\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("v=%v\n", v)
	}

	for i := range a {
		fmt.Printf("i=%v\n", i)
	}

}
