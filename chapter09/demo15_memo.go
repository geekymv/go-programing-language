package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func httpGetBodyV2(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// Func 函数类型
type FuncV2 func(key string) (interface{}, error)

type resultV2 struct {
	value interface{}
	err   error
}

// Memo 缓存调用 Func 的结果
type MemoV2 struct {
	f     FuncV2
	mu    sync.Mutex // 保护cache
	cache map[string]resultV2
}

func NewV2(f FuncV2) *MemoV2 {
	return &MemoV2{
		f:     f,
		cache: make(map[string]resultV2),
	}
}

func (m *MemoV2) GetV2(url string) (interface{}, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	res, ok := m.cache[url]
	if !ok {
		// 没有缓存，调用函数 Func
		res.value, res.err = m.f(url)
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
	}

	memo := NewV2(httpGetBodyV2)

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			start := time.Now()
			val, err := memo.GetV2(url)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("%s, %d, %d bytes\n", url, time.Since(start)/time.Millisecond, len(val.([]byte)))
		}(url)
	}
	wg.Wait()

}
