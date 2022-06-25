package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {

	var w io.Writer
	w = os.Stdout
	// 类型断言是一个作用在接口值上的操作，其中 x 是一个接口类型的表达式
	// x.(T) 如果 T 是具体类型，检查 x 的动态类型是否就是T
	f := w.(*os.File)
	fmt.Printf("%T\n", f)

	b, ok := w.(*bytes.Buffer) // 使用ok 可以防止 panic
	if ok {
		fmt.Printf("%T\n", b)
	}

	// x.(T) 如果 T 是接口类型，检查 x 的动态类型是否满足T（也就是说判断 x 的动态类型是否实现了接口 T）
	// 如果满足，返回的结果为接口类型 T
	rw := w.(io.ReadWriter)
	rw.Write([]byte("hello\n"))

}
