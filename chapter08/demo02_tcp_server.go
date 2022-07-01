package main

import (
	"fmt"
	"io"
	"log"
	"net"
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
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		// 发送当前时间给客户端
		_, err := io.WriteString(conn, time.Now().Format("2006-01-02 15:04:05\n"))
		if err != nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
}
