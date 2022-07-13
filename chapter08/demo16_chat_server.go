package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// 对外发送消息的通道
type client chan<- string

var (
	// 客户端连接通知通道
	entering = make(chan client)
	// 客户端断开通知通道
	leaving = make(chan client)
	// 全局消息通道
	messages = make(chan string)
)

func broadcaster() {
	// 所有连接的客户端
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			// 把接收到的消息广播给所有客户端
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			// 客户端连接
			clients[cli] = true
		case cli := <-leaving:
			// 客户端断开
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConnV2(conn net.Conn) {
	// 客户端连接对应的一个 channel
	clientChan := make(chan string)
	// 将 cli 通道中的消息发送给客户端
	go clientWriter(conn, clientChan)

	// 获取客户端连接地址信息
	who := conn.RemoteAddr().String()
	// 将客户端链接地址发送消息通道（服务端返回客户端的连接信息）
	clientChan <- "server send:" + who

	// 发送消息
	messages <- who + " has arrived"
	// 客户端连接
	entering <- clientChan

	// 循环读取客户端发送的消息
	input := bufio.NewScanner(conn)
	for input.Scan() {
		msg := input.Text()
		if msg == "close" {
			log.Println(who + " exist")
			break
		}
		messages <- who + ":" + msg
	}

	// 客户端离开了
	leaving <- clientChan
	messages <- who + " has left"

	// 关闭连接
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		// 将通道中的消息发送给客户端
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	port := 8000
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("server start success at %d", port)

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnV2(conn)
	}
}
