package main

import (
	"log"
	"os"
)

// 发起 GET 请求，解析 HTML，并返回 HTML 中存在的链接
func extract(url string) ([]string, error) {

	// mock data
	return []string{}, nil
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

			go func(link string) {
				worklist <- crawl(link)
			}(link)
		}
	}

}
