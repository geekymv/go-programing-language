package main

import "fmt"

func typeSwitch(x interface{}) {
	// 类型分支
	switch x := x.(type) {
	case int:
		fmt.Println("x=", x)
	case bool:
		fmt.Println("bool=", x)
	default:
		fmt.Println("default")
	}
}

func main() {
	//var x int = 1
	x := true
	typeSwitch(x)
}
