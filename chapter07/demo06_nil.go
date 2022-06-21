package main

import (
	"bytes"
	"io"
)

const debug = false

// 接口的动态类型是 *bytes.Buffer, 动态值nil
// 一个接口是否为nil取决于它的动态类型
func f(w io.Writer) {
	// 空指针的非空接口
	if w != nil { // 比较的是动态类型
		w.Write([]byte("done!")) // 调用nil的任何方法都会panic
	}
}

func main() {
	//var buf io.Writer
	var buf *bytes.Buffer // nil
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
}
