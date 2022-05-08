package main

import "fmt"

func main() {
	s := "hello,world"
	// len 返回字符串的字节数
	fmt.Println("len:", len(s)) // 11
	fmt.Println("s[1]", s[1])   // 101 'e'

	// fmt.Println("s[11]", s[11]) // 发生 panic

	// s[i:j] 包括i，不包括j。
	// i 默认值0，j 默认值是len(s)
	// i 和 j 可以省略其中一个或两个都省略，则取默认值
	fmt.Println("s[0]-s[5]", s[0:5])

	// 2 + 6 * 3
	s2 := "Go程序设计语言"
	fmt.Println("len:", len(s2)) // 20
	fmt.Println("s2[2]", s2[2])  // 231 'ç'

	// 程
	fmt.Println(string([]byte{s2[2], s2[3], s2[4]}))

	// 字符串不可变，字符串是不可变的字节序列

	// s2[0] = 'H' // 编译报错

	// 原生的字符串字面量 反引号
	s3 := `hello
		\r\n world`
	fmt.Println("s3:", s3)

}
