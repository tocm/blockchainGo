package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

const AES_KEY_ENCRYPT  = "1234567890123456" //16byte [128bit]

//AES对称加密，需要首先对明文补码
//abcdefghi
//abcdefghi77777777
//abcd4444

//补码
//PKCS5的分组是以8为单位
//PKCS7的分组长度为1-255
func PKCS7Padding(org []byte, blockSize int) []byte {
	pad:=blockSize-len(org)%blockSize
	padArr:=bytes.Repeat([]byte{byte(pad)},pad)
	return append(org,padArr...)
}

/**
	AES加密
	AES明文分组长度为128位，16字节。
	密钥长度为 16字节（128bit），24字节(192bit)，32字节(256bit)。
 */
func AESEncrypt(org []byte , key [] byte) []byte{
	//检验秘钥
	block,_:=aes.NewCipher(key)
	//对明文进行补码
	org = PKCS7Padding(org,block.BlockSize())
	//设置加密模式
	blockMode:=cipher.NewCBCEncrypter(block,key)

	//创建密文缓冲区
	cryted:=make([]byte,len(org))
	//加密
	blockMode.CryptBlocks(cryted,org)
	//返回密文
	return cryted

}