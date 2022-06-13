package main

import "fmt"

type ByteCounter int

// Write 实现了 io.Writer 接口
func (c *ByteCounter) Write(p []byte) (n int, err error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	fmt.Fprintf(&c, "hello %s", "tom")
	fmt.Println(c)
}
