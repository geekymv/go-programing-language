package main

import "fmt"

func main() {

	var rmdirs []func()

	tempDirs := []string{"/usr/local/a", "/usr/local/b", "/usr/local/c"}

	for i := 0; i < len(tempDirs); i++ {
		// 创建一个函数，延迟调用
		rmdirs = append(rmdirs, func() {
			// 这个 dir 会是最后一个迭代的dir变量值
			fmt.Println("i", i)
			//fmt.Println("remove dir", tempDirs[i])
		})
	}

	// 模拟循环执行rmdir函数
	for _, rmdir := range rmdirs {
		rmdir()
	}

}
