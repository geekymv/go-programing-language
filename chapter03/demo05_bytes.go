package main

import (
	"bytes"
	"fmt"
)

func main() {

	// 无需初始化，零值可用
	var buf bytes.Buffer

	buf.WriteString("hello")
	buf.WriteRune(',')
	buf.WriteByte('w')

	fmt.Println(buf.String())

}
