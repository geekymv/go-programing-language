package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func httpGetBodyV4(url string) (interface{}, error) {
	log.Println("url=", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// Func 函数类型
type FuncV4 func(key string) (interface{}, error)

type resultV4 struct {
	value interface{}
	err   error
}

type entryV4 struct {
	res   resultV4
	ready chan struct{} // res 准备好后，channel 会关闭
}

// Memo 缓存调用 Func 的结果
type MemoV4 struct {
	f     FuncV4
	mu    sync.RWMutex // 保护cache
	cache map[string]*entryV4
}

func NewV4(f FuncV4) *MemoV4 {
	return &MemoV4{
		f:     f,
		cache: make(map[string]*entryV4),
	}
}

// GetV4 RWLock + channel
func (m *MemoV4) GetV4(url string) (interface{}, error) {
	// 共享锁，可以并发读
	m.mu.RLock()
	e := m.cache[url]
	m.mu.RUnlock()
	if e == nil {
		// 排他锁
		m.mu.Lock()
		// double check
		e = m.cache[url]
		if e != nil {
			<-e.ready
		}

		// e 是 entry 指针
		e = &entryV4{
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
		"https://www.baidu.com",
		"https://www.baidu.com",
		"https://www.baidu.com",
		"https://www.baidu.com",
		"https://www.baidu.com",
	}

	memo := NewV4(httpGetBodyV4)

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			start := time.Now()
			val, err := memo.GetV4(url)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("%s, %d, %d bytes\n", url, time.Since(start)/time.Millisecond, len(val.([]byte)))
		}(url)
	}
	wg.Wait()

}
