package main

import (
	"bufio"
	"fmt"
	"net"
)

// tcp server demo

func process(conn net.Conn) {
	defer conn.Close() // 处理完之后要关闭这个连接
	// 针对当前的连接做数据的发送和接收操作
	for {
		reader := bufio.NewReader(conn) // 从该网络套接字中读取
		var buf [128]byte               // 应用层缓冲区
		n, err := reader.Read(buf[:])   // 读到这个切片中
		if err != nil {
			fmt.Printf("read from conn failed, err:%v\n", err)
			break
		}
		recvString := string(buf[:n])
		fmt.Println("接收到的数据：", recvString)
		conn.Write([]byte("ok")) // 把收到的数据返回给客户端
	}
}

func main() {
	// 1. 开启服务
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("listen failed, err:%v\n", err)
		return
	}
	for {
		// 2. 等待客户端来建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			continue
		}
		// 获取到了一个与客户端建立的连接
		// 3. 启动一个单独的goroutine去处理连接
		go process(conn)
	}

}
