package main

import (
	"net"
	"fmt"
	"blockchainGo/crypt/aes"
)

/**
	TCP 传输 模拟server side
 */
func main()  {

	listener, _ := net.Listen("tcp","localhost:5555")

	defer listener.Close() //当程序退出时确保关闭listener

	//处理接收数据
	for {
		new_conn, _ := listener.Accept();

		go recvMessage(new_conn)
	}

	fmt.Println("-----tcp socket finish----")
}

func recvMessage(conn net.Conn)  {
	recvBuf := make([]byte,1024);
	defer conn.Close()

	for{
		index,err := conn.Read(recvBuf)//读取数据到buf arrays

		if err != nil {
			fmt.Println("conn closed")
			return
		}
		//data[:n]就是接受到的密文
		fmt.Println("接收到client 字节数据为：",recvBuf[:index])
		fmt.Println("接收到client 原始字符加密后的数据为：",string(recvBuf[:index]))
		fmt.Println("从客户端接收到的名文为：",string(aes.AESDecrypt(recvBuf[:index],[]byte(aes.AES_KEY_ENCRYPT))))
	}
}