package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func httpGetBodyV5(url string) (interface{}, error) {
	log.Println("url=", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// Func 函数类型
type FuncV5 func(key string) (interface{}, error)

type resultV5 struct {
	value interface{}
	err   error
}

type entryV5 struct {
	res   resultV5
	ready chan struct{} // res 准备好后，channel 会关闭
}

func (e *entryV5) call(f FuncV5, url string) {
	// 发起请求，读取数据
	e.res.value, e.res.err = f(url)
	// 关闭 channel，广播
	close(e.ready)
}

func (e *entryV5) deliver(response chan<- resultV5) {
	// 等待数据读取完毕
	<-e.ready
	response <- e.res
}

// 封装了请求URL和响应
type request struct {
	url      string
	response chan<- resultV5
}

type MemoV5 struct {
	// 接受请求的 channel
	requests chan request
}

func NewV5(f FuncV5) *MemoV5 {
	m := &MemoV5{
		requests: make(chan request),
	}
	go m.server(f)
	return m
}

func (m *MemoV5) GetV5(url string) (interface{}, error) {
	response := make(chan resultV5)
	// 封装请求，并将请求发送给 channel
	m.requests <- request{
		url:      url,
		response: response,
	}
	// 阻塞等待结果
	res := <-response
	return res.value, res.err
}

// 只有一个 goroutine 操作 server
func (m *MemoV5) server(f FuncV5) {
	// 只有一个 goroutine 访问 map
	cache := make(map[string]*entryV5)

	fmt.Println("for start")
	for req := range m.requests {
		e := cache[req.url]
		if e == nil {
			e = &entryV5{
				ready: make(chan struct{}),
			}
			cache[req.url] = e
			go e.call(f, req.url)
		}
		go e.deliver(req.response)
	}
	fmt.Println("for end")
}

func main() {

	urls := []string{
		"https://www.baidu.com",
		"https://www.hao123.com",
		"https://www.baidu.com",
		"https://www.baidu.com",
		"https://www.baidu.com",
		"https://www.baidu.com",
		"https://www.baidu.com",
		"https://www.hao123.com",
	}

	memo := NewV5(httpGetBodyV5)

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			start := time.Now()
			val, err := memo.GetV5(url)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("%s, %d, %d bytes\n", url, time.Since(start)/time.Millisecond, len(val.([]byte)))
		}(url)
	}
	wg.Wait()

}
