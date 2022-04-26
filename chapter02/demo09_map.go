package main

import (
	"fmt"
	"strconv"
)

// ID 类型声明 定义一个新的命名类型
type ID uint32

func (i ID) String() string {
	return strconv.Itoa(int(i))
}

func main() {

	m := make(map[string]int)
	m["a"] = 1

	// ok 判断 key 是否存在
	v, ok := m["b"]
	if ok {
		fmt.Println("v=", v)
	}

	var id ID = 100
	fmt.Println("id=", id.String())

}

func sayHello() {

}
