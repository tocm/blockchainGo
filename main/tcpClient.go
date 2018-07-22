package main

import (
	"net"
	"fmt"
	"blockchainGo/crypt/aes"
)

/**
	TCP client side
 */

func main()  {
	//广域网，UDP+NAT

	//构建服务器连接
	conn,_:=net.ResolveTCPAddr("tcp","127.0.0.1:5555")

	//连接拨号
	n,_:=net.DialTCP("tcp",nil,conn)
	//发送数据

	n.Write(testSendEncryptMsg())
	fmt.Println("发送结束")

}

func testSendEncryptMsg() []byte  {

	msg := "Hello, test the AES Encrypt and Tcp socket"
	return aes.AESEncrypt([]byte(msg), []byte(aes.AES_KEY_ENCRYPT)) //aes key 16byte
}
