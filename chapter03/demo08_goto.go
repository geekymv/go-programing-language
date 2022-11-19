package main

import "fmt"

func exit(i int) {
	if i == 1 {
		fmt.Println("aaa")
		goto exit
	}
	fmt.Println("bbb")

exit:
	// 无论是否执行了 goto 下面这行都会执行
	fmt.Println("exit")
}

func main() {

	exit(1)

}
