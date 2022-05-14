package main

import "fmt"

func main() {

	// slice
	//var s1 []string = []string{"a", "b", "c"}

	// array
	//var s2 [3]string = [3]string{"a", "b", "c"}

	// 指向数组的 slice
	s3 := []string{"a", "b", "c"}
	fmt.Printf("len(s3)=%d, cap(s3)=%d\n", len(s3), cap(s3))

	//s4 := [3]string{"a", "b", "c"}

	var s5 []string
	s6 := []string{}
	fmt.Println("s5", s5 == nil)
	fmt.Println("len(s5)", len(s5))
	fmt.Println("s6", s6 == nil)
	fmt.Println("len(s6)", len(s6))

	// 判断一个slice是否为空 len(s) == 0

}
