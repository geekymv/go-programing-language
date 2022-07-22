package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func httpGetBody(url string) (interface{}, error) {
	fmt.Println("url=", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// Func 函数类型
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// Memo 缓存调用 Func 的结果
type Memo struct {
	f     Func
	cache map[string]result
}

func New(f Func) *Memo {
	return &Memo{
		f:     f,
		cache: make(map[string]result),
	}
}

// Get 先从缓存中取，如果获取不到，则调用 Func 函数，发起请求，最后缓存结果
func (m *Memo) Get(url string) (interface{}, error) {
	res, ok := m.cache[url]
	if !ok {
		// 想要实现的效果是：同样的 url 只会访问一次
		// 没有缓存，调用函数 Func
		res.value, res.err = m.f(url)
		// concurrent map writes
		m.cache[url] = res
	}
	return res.value, res.err
}

func main() {

	urls := []string{
		"https://www.baidu.com",
		"https://www.baidu.com",
		"https://www.baidu.com",
		"https://www.baidu.com",
		"https://www.baidu.com",
		"https://www.baidu.com",
		"https://www.baidu.com",
		"https://www.baidu.com",
		"https://www.baidu.com",
	}

	memo := New(httpGetBody)

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			start := time.Now()

			val, err := memo.Get(url)

			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("%s, %d, %d bytes\n", url, time.Since(start)/time.Millisecond, len(val.([]byte)))
		}(url)
	}
	wg.Wait()

}
