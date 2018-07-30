package main

import (
	"net"
	"fmt"
	"blockchainGo/crypt/rsa"
	appUtils "blockchainGo/utils"
	"blockchainGo/crypt/aes"
)

/**
	TCP 传输 模拟server side

实例目标：
1. 使用AES对称私钥对数据进行解密
2. 使用非对称RSA公钥对数据进行验证签名
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

/**
	协议： [标识位(1byte)签名数据长度 + 签名数据 + 文本数据]
 */
func recvMessage(conn net.Conn)  {
	fmt.Println("--------Begin recvMessage--------")
	var tempSize int = 256
	dataSlice := make([]byte, tempSize);
	recvBuf := make([]byte,tempSize);
	defer conn.Close()

	var index int =0
	for{
		n,err := conn.Read(recvBuf)//读取数据到buf arrays
		fmt.Println("index  : ",index)
		fmt.Println("conn new read n from connection io : ",n)
		if err != nil {
			fmt.Println("read end conn closed")
			parseMessage(dataSlice[:index])
			//log.Panic(err)
			return
		}
		//data[:n]就是接受到的密文
		fmt.Println("接收到client 字节数据为：",recvBuf[:n])
	//	fmt.Println("接收到client 原始字符加密后的数据为：",string(recvBuf[:n]))

		sum := n + index

		fmt.Printf("sum = %d, dataslice = %d \n",sum , len(dataSlice))
		if sum > len(dataSlice) {
			//扩容
			new_dataSlice := make([]byte, len(dataSlice) *  2)
			//拷贝原来的数据到新数组
			copy(new_dataSlice,dataSlice)
			//更新新数组
			dataSlice = new_dataSlice
			fmt.Println("-------------> 扩容 new size ",len(new_dataSlice))
		}

		fmt.Printf("index = %d, sum = %d, recv n = %d, lenDataSlice = %d \n",index, sum, n, len(dataSlice))
		copy(dataSlice[index:sum],recvBuf[:n])

		index = sum

	}


}

/**
	解析服务器传送过来的数据
	recvBuf[] byte 实际数据大小的数组
 */
func parseMessage(recvBuf[] byte)  {
	fmt.Println("==============>parseMessage begin---")
	fmt.Println("解密前字节长度： ",len(recvBuf))
	//AES 对称解密
	deCryptoData := aes.AESDecrypt(recvBuf,[]byte(aes.AES_KEY_ENCRYPT))

	fmt.Println("解密后原文字节长度：", len(deCryptoData))

	//分离解密后的原文数据，根据[标识位int(4byte)签名数据长度 + 签名数据 + 文本数据]
	startPrefix := 4
	lenSignBytes := deCryptoData[0:startPrefix] //直接取出标识位-签名数据长度
	fmt.Println("解密后签名标识字节长度数组：", lenSignBytes)
	//byte 转成 int 型
	//lenSignInt := lenSignBytes[3] & 0xff | (lenSignBytes[2] & 0xff) << 8 | (lenSignBytes[1] & 0xff) << 16 | (lenSignBytes[0] & 0xff) << 24;

	lenSignInt := appUtils.BytesToInt32Bigendian(lenSignBytes);
	fmt.Println("解密后签名数据长度：", lenSignInt)
	startSignData := startPrefix;
	endSignData := 4 + lenSignInt
	fmt.Printf("startSignData %d, endsignData =%d, deCryptoData=%d \n",startSignData,endSignData,len(deCryptoData))
	//取出真实签名的数据
	signData := deCryptoData[startSignData : endSignData]
	orgContent := deCryptoData[4 + lenSignInt :]

	fmt.Println("解密后", deCryptoData)
	fmt.Println("加密后-sign:", signData)

	//验证签名
	isSignVarifySuccess := rsa.RSAVerifySign(signData)

	if isSignVarifySuccess {
		fmt.Println("从客户端接收到的原始名文为：",string(orgContent))
	}
}