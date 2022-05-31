package main

import "fmt"

func f(x int) {
	fmt.Printf("f(%d)=%d\n", x, 0/x) // x = 0 panic
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func main() {
	//f(3)
	defer func() {
		// 恢复
		err := recover()
		fmt.Println("recover err:", err)
	}()

	for i := 0; i < 10; i++ {
		if i == 5 {
			// panic 程序退出
			panic("i == 5")
		}
		fmt.Println("i=", i)
	}

}
