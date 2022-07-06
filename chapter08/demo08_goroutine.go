package main

import (
	"fmt"
	"sync"
	"time"
)

func makeImage(filename string) {
	time.Sleep(1 * time.Second)
	fmt.Println("make image = ", filename)
}

/*
func ImageFile(fileNames []string) {
	// 等待所有goroutine执行完再返回
	ch := make(chan struct{})
	for _, f := range fileNames {
		// f 变量重用
		go func(f string) {
			makeImage(f)
			ch <- struct{}{}
		}(f)
	}

	for range fileNames {
		// 会阻塞
		<-ch
	}
}
*/

func ImageFile(fileNames []string) {
	// 等待所有goroutine执行完再返回
	var wg sync.WaitGroup
	for _, f := range fileNames {
		// f 变量重用
		// Add 说明有一个goroutine需要等待
		wg.Add(1)
		go func(f string) {
			// Done 表示一个goroutine执行完成
			defer wg.Done()
			makeImage(f)
		}(f)
	}

	// 等待所有goroutine执行完
	wg.Wait()

}

func main() {
	fileNames := []string{"a.jpg", "b.jpg", "c.jpg", "d.jpg", "e.jpg", "f.jpg"}
	ImageFile(fileNames)
}
