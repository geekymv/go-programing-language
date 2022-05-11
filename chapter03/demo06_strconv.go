package main

import (
	"fmt"
	"strconv"
)

func main() {

	i := 123
	// 数字转字符串
	s := fmt.Sprintf("%d", i)
	fmt.Printf("type:%T, s:%v\n", s, s)

	s = strconv.Itoa(i)
	fmt.Printf("type:%T, s:%v\n", s, s)

	// 二进制形式的字符串
	s = strconv.FormatInt(int64(i), 2)
	fmt.Printf("type:%T, s:%v\n", s, s)

	s = fmt.Sprintf("%b", i)
	fmt.Printf("type:%T, s:%v\n", s, s)

	fmt.Println("---------------------------")
	// 字符串转整数
	i, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Printf("type:%T, s:%v\n", i, i)
	}

	i64, err := strconv.ParseInt("123", 10, 0)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Printf("type:%T, s:%v\n", i64, i64)
	}

}
