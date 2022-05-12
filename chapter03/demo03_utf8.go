package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s := "Hello,世界"
	// 字节个数
	fmt.Println(len(s)) // 12
	// 码点个数
	fmt.Println(utf8.RuneCountInString(s)) // 8

	// DecodeRuneInString 返回第一个UTF-8编码的码点rune 和 第一个rune字节长度
	u, size := utf8.DecodeRuneInString(s)
	fmt.Printf("u=%c\n", u)
	fmt.Println("size=", size)

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("r:%c\tsize:%d\n", r, size)
		// 定位下一个文字符号
		i += size
	}

	fmt.Println("-----------------------------")

	// for range 遍历字符串 UTF-8 隐式解码
	for i, r := range s {
		// i 是 r 码点占用字节的起始下标
		fmt.Print(i)
		fmt.Printf("\t")
		fmt.Printf("%c\n", r)
	}

	fmt.Println("-----------------------------")

	// 将字符串转成 Unicode 码点序列
	r := []rune(s) // int32 存放的是 Unicode 码点
	fmt.Println("len", len(r))
	for i, c := range r {
		// i 是码点下标
		fmt.Print(i)
		fmt.Printf("\t")
		fmt.Printf("%c\n", c)
	}

	fmt.Println("r=", string(r))

}
