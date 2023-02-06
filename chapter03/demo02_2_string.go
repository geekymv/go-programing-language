package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "hello,世界"
	fmt.Println("len(hello,世界)=", utf8.RuneCountInString(str))
}
