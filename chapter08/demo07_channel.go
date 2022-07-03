package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string, 3)

	go func() {
		ch <- request("https://a.com")
	}()

	go func() {
		ch <- request("https://b.com")
	}()

	go func() {
		ch <- request("https://c.com")
	}()

	// channel 为空，这里会阻塞
	fmt.Println(<-ch)
}

func request(url string) string {
	switch url {
	case "https://a.com":
		time.Sleep(2 * time.Second)
		return "a"
	case "https://b.com":
		time.Sleep(1 * time.Second)
		return "b"
	case "https://c.com":
		time.Sleep(3 * time.Second)
		return "c"
	}
	return "--"
}
