package main

import "fmt"

// 导入路径
import "go-programing-language/util"

func init() {
	fmt.Println("init invoked")
}

func main() {
	fmt.Println("main invoked")

	str := strutil.StrToUpper("hello")
	fmt.Println("str=", str)
}
