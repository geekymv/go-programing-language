package main

import "fmt"

// go1.16.14 版本
// runtime/slice.go
// growslice 方法
// https://mp.weixin.qq.com/s/p572g5KcSwy2ri40d1cPTg
func main() {
	// cap 小于1024
	// newcap = oldcap * 2
	s1 := make([]int, 0, 1023)
	for i := 0; i < 1024; i++ {
		s1 = append(s1, i)
	}
	fmt.Println(len(s1), cap(s1)) // 1024 2048

	s := make([]int, 0, 1024)

	for i := 0; i < 1024; i++ {
		s = append(s, i)
	}
	fmt.Println(len(s), cap(s)) // 1024 1024
	// cap 容量大于1024
	// newcap = oldcap + oldcap / 4
	s = append(s, 1)
	fmt.Println(len(s), cap(s)) // 1025 1280

}
