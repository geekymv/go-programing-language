package main

import (
	"errors"
	"log"
	"net"
	"os"
	"syscall"
	"time"
)

// https://gosamples.dev/broken-pipe/
/**
sudo tcpdump -i any -s 0 -B 524288 port 9090 -w ~/Desktop/DumpFile01.pcap

-i interface
-B buffer size (unit Kib)
-s snaplen
https://hackertarget.com/tcpdump-examples/
*/
func server() {
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
	data := make([]byte, 5)
	if _, err := conn.Read(data); err != nil {
		log.Fatal("server", err)
	}
	// 读取一次数据，关闭连接
	conn.Close()
}

func client() {
	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		log.Fatal("client", err)
	}

	log.Println("write to make the connection closed on the server side")
	if _, err := conn.Write([]byte("hello")); err != nil {
		log.Printf("client: %v", err)
	}

	time.Sleep(1 * time.Second)

	// 服务端已经关闭了连接，会返回 RST 包
	log.Println("write to generate an RST packet")
	if _, err := conn.Write([]byte("b")); err != nil {
		log.Printf("client: %v", err)
	}

	time.Sleep(1 * time.Second)

	// 已经收到了服务端已经返回 RST 包，再发送数据就会报错 write: broken pipe
	log.Println("write to generate the broken pipe error")
	if _, err := conn.Write([]byte("c")); err != nil {
		log.Printf("client: %v", err)
		if errors.Is(err, syscall.EPIPE) {
			log.Print("This is broken pipe error")
		}
	}
}

func main() {
	go server()

	time.Sleep(10 * time.Second) // wait for server to run

	client()
}
