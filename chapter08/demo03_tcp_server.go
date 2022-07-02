package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		// 会阻塞，接收客户端的连接
		fmt.Println("accept before")
		conn, err := listener.Accept()
		fmt.Println("accept after", conn.RemoteAddr().String())
		if err != nil {
			log.Print(err)
			continue
		}
		// 服务端需要并发处理多个客户端的连接
		// 开启一个 goroutine 处理一个连接
		go handleConn2(conn)
	}
}

func handleConn2(conn net.Conn) {
	defer conn.Close()
	// 读取客户端输入的信息
	input := bufio.NewScanner(conn)
	for input.Scan() {
		// echo 方法会sleep 方法没有执行完 不会执行下一次的for循环，也就是说不会读取下一个客户端的输入
		// 新开启一个 goroutine 处理客户端的每次输入
		go echo(conn, input.Text(), 2*time.Second)
	}
}

func echo(conn net.Conn, info string, delay time.Duration) {
	// 大写
	fmt.Fprintln(conn, "\tresp:", strings.ToUpper(info))
	time.Sleep(delay)

	// 原样返回
	fmt.Fprintln(conn, "\tresp:", info)
	time.Sleep(delay)

	// 小写
	fmt.Fprintln(conn, "\tresp:", strings.ToLower(info))
}
