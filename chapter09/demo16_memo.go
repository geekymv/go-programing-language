package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func httpGetBodyV3(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// Func 函数类型
type FuncV3 func(key string) (interface{}, error)

type resultV3 struct {
	value interface{}
	err   error
}

type entry struct {
	res   resultV3
	ready chan struct{} // res 准备好后，channel 会关闭
}

// Memo 缓存调用 Func 的结果
type MemoV3 struct {
	f     FuncV3
	mu    sync.Mutex // 保护cache
	cache map[string]*entry
}

func NewV3(f FuncV3) *MemoV3 {
	return &MemoV3{
		f:     f,
		cache: make(map[string]*entry),
	}
}

func (m *MemoV3) GetV3(url string) (interface{}, error) {
	// 排他锁
	m.mu.Lock()
	e := m.cache[url]
	if e == nil {
		// e 是 entry 指针
		e = &entry{
			ready: make(chan struct{}),
		}
		// 没有准备好数据 res 的 entry
		m.cache[url] = e
		m.mu.Unlock()

		// 发起请求，读取数据
		e.res.value, e.res.err = m.f(url)

		// 关闭 channel，广播
		close(e.ready)

	} else {
		m.mu.Unlock()
		// 等待数据读取完毕
		<-e.ready
	}
	return e.res.value, e.res.err
}

func main() {

	urls := []string{
		"https://www.baidu.com",
		"https://www.baidu.com",
		"https://www.baidu.com",
		"https://www.baidu.com",
	}

	memo := NewV3(httpGetBodyV3)

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			start := time.Now()
			val, err := memo.GetV3(url)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("%s, %d, %d bytes\n", url, time.Since(start)/time.Millisecond, len(val.([]byte)))
		}(url)
	}
	wg.Wait()

}
