package main

import (
	"net"
	"fmt"
	"blockchainGo/crypt/aes"
	"blockchainGo/crypt/rsa"
	appUtils "blockchainGo/utils"
)


/**
	TCP client side

实例目标：
1. 使用非对称RSA私钥对数据进行签名
2. 使用AES对称公钥对数据进行加密后传输
 */

func main()  {
	//广域网，UDP+NAT
	//构建服务器连接
	conn,_:=net.ResolveTCPAddr("tcp","127.0.0.1:5555")

	//连接拨号
	n,_:=net.DialTCP("tcp",nil,conn)
	//发送数据

	n.Write(part1_EncryptoAndRsaSign())
	fmt.Println("发送结束")

}

/**
	发送协议： [标识位(1byte)签名数据长度 + 签名数据 + 文本数据]
 */
func part1_EncryptoAndRsaSign() []byte  {

	fmt.Println("part1:------ 加密及RSA签名 demo --------")
	var sendData []byte

	//【RSA 签名】
	sign := rsa.RSASign()

	//要发送的文本
	contentTxt := "Hello, test the AES Encrypt and Tcp socket "
	//签名文件的字节长度
	lenSign := len(sign)

	//整型转成字节数组
	lenSignBytes := appUtils.IntToBytesBigendian(int32(lenSign))

	//重新拼接数据到发送数组[标识位int32(4byte)签名数据长度 + 签名数据 + 文本数据]
	sendData = make([]byte, len(lenSignBytes) + lenSign +len(contentTxt))

	fmt.Printf("加密前各字节长度 (len(lenSignBytes) %d + lenSign %d + len(content)%d) = %d \n",len(lenSignBytes),lenSign, len(contentTxt),len(sendData))
	copy(sendData[0:len(lenSignBytes)],lenSignBytes)
	copy(sendData[len(lenSignBytes): len(lenSignBytes) + lenSign], []byte(sign))
	copy(sendData[len(lenSignBytes) + lenSign:],contentTxt)

	fmt.Println("加密前-signLen:", lenSignBytes)
	fmt.Println("加密前-sign:", sign)
	fmt.Println("加密前", sendData)


	/*
	使用【AES】对称进行传送数据加密
	由于RSA已经在数据中用了RSA签名，根据RSA中加密的约束条件：传送加密的明文必须要少于生成私钥时的字节大小，
	所以如果在传输的数据内已经使用了RSA数据签名，即数据内已经包含了签名的字节长度，这样的话传输的数据就永远大于生成私钥时的字节大小，这样就会导致RSA加密失败，
	所以如果已经用了非对称RSA签名的数据，只能使用对称的加密方式
	*/
	fmt.Println("use AES encrypto")
	entryptoBytes := aes.AESEncrypt(sendData, []byte(aes.AES_KEY_ENCRYPT)) //aes key 16byte
	fmt.Printf("发送前--加密后长度 %d \n",len(entryptoBytes))
	return entryptoBytes
}


