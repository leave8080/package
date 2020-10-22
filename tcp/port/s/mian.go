//Server.go
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	//服务端在本机8888端口建立tcp监听

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	listener.Addr()
	fmt.Println(listener.Addr())
	ServerHandleError(err, "net.listen")

	for {
		//循环接入所有客户端得到专线连接
		conn, e := listener.Accept()
		ServerHandleError(e, "listener.accept")
		//开辟独立协程与该客聊天
		go ChatWith(conn)
	}

}

func ServerHandleError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}

//在conn网络专线中与客户端对话
func ChatWith(conn net.Conn) {

	//创建消息缓冲区
	buffer := make([]byte, 1024)
	for {
		///---一个完整的消息回合

		//读取客户端发来的消息放入缓冲区
		n, err := conn.Read(buffer)
		ServerHandleError(err, "conn.read buffer")

		//转化为字符串输出
		clientMsg := string(buffer[0:n])
		fmt.Printf("收到消息", conn.RemoteAddr(), clientMsg)

		//回复客户端消息
		if clientMsg != "im off" {
			conn.Write([]byte("已读:" + clientMsg))
		} else {
			conn.Write([]byte("bye"))
			break
		}
	}
	conn.Close()
	fmt.Printf("客户端断开连接", conn.RemoteAddr())

}
