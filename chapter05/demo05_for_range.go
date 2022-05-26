package main

import "fmt"

func main() {

	var rmdirs []func()

	tempDirs := []string{"/usr/local/a", "/usr/local/b", "/usr/local/c"}

	// dir 不会每次创建一个变量，而是将每次迭代的值赋值给dir
	// dir 是同一个地址
	for _, dir := range tempDirs {
		fmt.Printf("mk dir:%v, dir:%p\n", dir, &dir)

		// 解决方式 声明一个变量d
		d := dir
		// 创建一个函数，延迟调用
		rmdirs = append(rmdirs, func() {
			// 这个 dir 会是最后一个迭代的dir变量值
			fmt.Println("remove dir", d)
		})
	}

	// 模拟循环执行rmdir函数
	for _, rmdir := range rmdirs {
		rmdir()
	}

}
