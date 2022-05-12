package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "/a/b/c.go"

	// 从后往前
	i := strings.LastIndex(s, "/")
	fmt.Println("i=", i)

	s = s[i+1:]
	fmt.Println("s=", s)

}
