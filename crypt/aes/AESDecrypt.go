package aes

import (
	"crypto/aes"
	"crypto/cipher"
)

//去码
func PKCS7UnPadding(org []byte )[]byte {
	//abcd4444
	l:=len(org)
	pad:=org[l-1]
	//org[0:4]
	return org[:l-int(pad)]
}


/**
	AES 解码
	密钥长度为 16字节（128bit），24字节(192bit)，32字节(256bit)。
 */
func AESDecrypt(cipherTxt []byte ,key []byte )[]byte {
	block,_:=aes.NewCipher(key)
	blockMode:=cipher.NewCBCDecrypter(block,key)
	//创建明文缓存
	org:=make([]byte,len(cipherTxt))
	//开始解密
	blockMode.CryptBlocks(org,cipherTxt)

	//去码
	org = PKCS7UnPadding(org)
	//返回明文
	return org

}

