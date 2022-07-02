package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 新开启一个 goroutine 读取服务端返回的信息，输出到控制台
	go mustCopy(os.Stdout, conn)

	// 读取控制台输入的信息发给服务端
	mustCopy(conn, os.Stdin)

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
