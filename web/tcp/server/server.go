package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("服务端读取消息失败,错误:%s\n", err)
			break
		}

		recvStr := string(buf[:n])
		fmt.Printf("收到客户端消息:%s\n", recvStr)
		conn.Write([]byte(recvStr))
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("监听客户端错误,错误:%s\n", err)
		return
	}
	fmt.Println("服务端启动成功,等待连接")
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("与客户端建立连接失败:%s\n", err)
			continue
		}
		fmt.Println("连接成功")
		go process(conn)
	}
}
