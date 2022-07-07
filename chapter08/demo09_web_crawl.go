package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// 发起 GET 请求，解析 HTML，并返回 HTML 中存在的链接
func extract(url string) ([]string, error) {
	fmt.Printf("url=%s\n", url)
	time.Sleep(1 * time.Second)
	// mock data
	return []string{"https://a.com", "https://b.com", "https://c.com"}, nil
}

func crawl(url string) []string {
	links, err := extract(url)
	if err != nil {
		log.Println(err)
	}

	return links
}

func main() {
	worklist := make(chan []string)
	// n 记录 channel 中元素个数，解决 deaklock 问题
	var n int

	n++
	go func() {
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if seen[link] {
				continue
			}
			// 设置已访问标志
			seen[link] = true
			// 访问一个url，会向 worklist 放入一个元素
			n++
			go func(link string) {
				worklist <- crawl(link)
			}(link)
		}
		// 取出一个元素
		n--
		// 判断 worklist channel 是否为空
		if n == 0 {
			break
		}
	}

}
