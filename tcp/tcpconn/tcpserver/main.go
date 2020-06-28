package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("listen err", err.Error())
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept err", err.Error())
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buff [128]byte
		n, err := reader.Read(buff[:])
		if err != nil {
			fmt.Println("read err", err.Error())
			break
		}
		fmt.Println("recvstr:", string(buff[:n]))
		conn.Write([]byte(string(buff[:n])))
	}
}
