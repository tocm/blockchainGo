package des

import (
	"fmt"
	"encoding/base64"
)


//DES加密中秘钥长度必须为8字节，3DES秘钥长度必须为24，AES加密算法中秘钥长度必须为16或24或32字节

const DES_CRYPT_KEY  =  "12345678" //DES 中密解必须是8位长度


/**
	对称加密算法
	加密与解密使用相同的key， 加密和解密是逆的
 */
func DesCryptMain()  {

	fmt.Println("-----DES 对称加密算法测试------- ")

	//补码
	padByte := PKCS5Padding([]byte("abc"),5)
	fmt.Println("after padding ", padByte)

	orgTxt := PKCS5UnPadding(padByte)
	fmt.Println("orginal text ",orgTxt)


	enCryptData := DESEnCrypt([]byte("Andy"), []byte(DES_CRYPT_KEY))//key 必须为8byte长度
	fmt.Println("加密后数组字节：",enCryptData)
	fmt.Println("加密后密文以base64 格式显示：",base64.StdEncoding.EncodeToString(enCryptData))

	deCryData := DESDeCrypt(enCryptData, []byte(DES_CRYPT_KEY))
	fmt.Println("解密后数组字节：",deCryData)
	fmt.Println("解密后字符串原文：",string(deCryData))
}