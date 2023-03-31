package main

import (
	"errors"
	"log"
	"net"
	"os"
	"syscall"
	"time"
)

// 当服务端关闭连接后，发送 RST 包给客户端之后，客户端从连接读数据，就会的到 connection reset by peer 错误
// 当服务端关闭连接后，发送 RST 包给客户端之后，客户端向连接写数据，就会的到 broken pipe 错误

func server2() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("server", err)
		os.Exit(1)
	}
	data := make([]byte, 1)
	if _, err := conn.Read(data); err != nil {
		log.Fatal("server", err)
	}

	conn.Close()
}

func client2() {
	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		log.Fatal("client", err)
	}

	if _, err := conn.Write([]byte("ab")); err != nil {
		log.Printf("client: %v", err)
	}

	time.Sleep(1 * time.Second) // wait for close on the server side

	data := make([]byte, 1)
	if _, err := conn.Read(data); err != nil {
		log.Printf("client: %v", err)
		if errors.Is(err, syscall.ECONNRESET) {
			log.Print("This is connection reset by peer error")
		}
	}
}

func main() {
	go server2()

	time.Sleep(3 * time.Second) // wait for server to run

	client2()
}
