package main

import (
	"encoding/binary"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:4150")
	if err != nil {
		return
	}
	conn.Write([]byte("  V2"))

	conn.Write([]byte("PUB hello\n"))

	//binary.Write(conn, binary.BigEndian, int32(5))
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, uint32(5))
	conn.Write(bs)

	conn.Write([]byte("worldhaha"))

	time.Sleep(2 * time.Minute)

	conn.Close()
}
