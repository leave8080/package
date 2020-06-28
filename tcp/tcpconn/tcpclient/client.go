package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("net dial err", err.Error())
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		inpuInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inpuInfo) == "Q" {
			return
		}
		_, err = conn.Write([]byte(inpuInfo))
		if err != nil {
			fmt.Println("writer err", err.Error())
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("read err", err.Error())
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
