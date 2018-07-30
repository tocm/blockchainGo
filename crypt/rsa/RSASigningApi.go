package rsa

import (
	"crypto/rsa"
	"crypto/rand"
	"crypto/md5"
	"fmt"
	"encoding/base64"
	"crypto"
	"log"
)

//--------------RSA 非对称加密运算 ----------------//
//---------直接使用go库api 生成private key and public key------//

//用公钥加密，私钥解密
//用私钥签名，公钥验证
//公钥是公开的，任何人可以使用公钥，私钥非公开（保存好）

const RSA_Sigining_Author  = "Andy.chen" //签名人

//使用go api 生成公钥，私钥
func RSAGenerateCryptKeyByApi() (privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey){
	//直接go api生成私钥
	priKey,_ :=rsa.GenerateKey(rand.Reader, 1024)

	//使用私钥创建公钥
	pubKey := priKey.PublicKey
	return priKey,&pubKey
}

//用公钥加密
func RSAEnCryptByApi(pubKey *rsa.PublicKey, sendData []byte) []byte {

	//使用公钥进行加密
	cipherData,_:=rsa.EncryptOAEP(md5.New(),rand.Reader,pubKey,sendData,nil)
	//打印密文
	fmt.Println("加密后密文为：",cipherData)
	fmt.Println("加密后密文转Base64为：",base64.StdEncoding.EncodeToString(cipherData))
	return cipherData
}

//用私钥解密
func RSADeCryptByApi(priKey *rsa.PrivateKey, cipherData []byte)  {

	//用私钥解密
	plaintext,_:=rsa.DecryptOAEP(md5.New(),rand.Reader,priKey, cipherData,nil)
	//打印解密后的结果
	fmt.Println("解密后明文为：",string(plaintext))

}

/**
	签名编程实现，私钥签名
 */
func RSAEnSigining(privKey *rsa.PrivateKey) []byte  {

	//实现hash散列
	h:=md5.New()
	h.Write([]byte(RSA_Sigining_Author))//签名人
	hashed:=h.Sum(nil)

	//通过RSA实现数字签名
	//数组签名的作用为验证是否被篡改，A->B,当B收到数据时，验证是否由A传递的消息

	opts:=rsa.PSSOptions{rsa.PSSSaltLengthAuto,crypto.MD5}

	//生成签名
	sig,_:=rsa.SignPSS(rand.Reader,privKey,crypto.MD5,hashed,&opts)

	fmt.Println("签名的结果为",sig)
	fmt.Println("加密后密文转Base64为：",base64.StdEncoding.EncodeToString(sig))

	return sig
}

/**
	验证签名编程实现，公钥验证签名的过程
 */
func RSAVerifySigining(pubKey *rsa.PublicKey, sigiData[]byte)  {

	//实现hash散列
	h:=md5.New()
	h.Write([]byte(RSA_Sigining_Author))//验证签名人
	hashed:=h.Sum(nil)

	opts:=rsa.PSSOptions{rsa.PSSSaltLengthAuto,crypto.MD5}

	//通过公钥实现验证签名
	err:=rsa.VerifyPSS(pubKey,crypto.MD5,hashed,sigiData,&opts)
	if err ==nil {
		fmt.Println("签名验证成功")
	} else {
		log.Panic(err)
	}

}